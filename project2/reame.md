```bash
# Initialization Commands for All Apps (Angular 17)
# Install Angular CLI 17 globally (if not installed)
npm install -g @angular/cli@17

# Create workspace and shell app (standalone by default in Angular 17)
ng new shell-app --routing --style=scss --standalone
cd shell-app
ng serve # Verify shell app runs

# Initialize Module Federation for shell (host)
ng add @angular-architects/module-federation --project shell-app --port 4200 --type host

# Create micro frontend apps in separate directories (standalone)
cd ..
ng new product-listing --style=scss --standalone
ng new user-auth --style=scss --standalone
ng new order-history --style=scss --standalone

# Initialize Module Federation for each remote
cd product-listing
ng add @angular-architects/module-federation --project product-listing --port 4201 --type remote
cd ../user-auth
ng add @angular-architects/module-federation --project user-auth --port 4202 --type remote
cd ../order-history
ng add @angular-architects/module-federation --project order-history --port 4203 --type remote
```

```javascript
// webpack.config.js (shell-app)
const { shareAll, withModuleFederationPlugin } = require('@angular-architects/module-federation/webpack');

module.exports = withModuleFederationPlugin({
  remotes: {
    'productListing': 'productListing@http://localhost:4201/remoteEntry.js',
    'userAuth': 'userAuth@http://localhost:4202/remoteEntry.js',
    'orderHistory': 'orderHistory@http://localhost:4203/remoteEntry.js'
  },
  shared: {
    ...shareAll({ singleton: true, strictVersion: true, requiredVersion: 'auto' })
  }
});
```

```typescript
// src/app/app.routes.ts (shell-app)
import { Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'products',
    loadChildren: () => import('productListing/routes').then(m => m.ProductListingRoutes)
  },
  {
    path: 'auth',
    loadChildren: () => import('userAuth/routes').then(m => m.UserAuthRoutes)
  },
  {
    path: 'orders',
    loadChildren: () => import('orderHistory/routes').then(m => m.OrderHistoryRoutes)
  },
  { path: '', redirectTo: '/products', pathMatch: 'full' }
];

export const appRoutes = routes;
```

```typescript
// src/main.ts (shell-app)
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter } from '@angular/router';
import { AppComponent } from './app/app.component';
import { appRoutes } from './app/app.routes';

bootstrapApplication(AppComponent, {
  providers: [
    provideRouter(appRoutes)
  ]
}).catch(err => console.error(err));
```

```html
<!-- src/app/app.component.html (shell-app) -->
<header class="bg-blue-600 text-white p-4">
  <nav class="container mx-auto flex justify-between items-center">
    <h1 class="text-2xl font-bold">E-Commerce</h1>
    <ul class="flex space-x-4">
      <li><a routerLink="/products" class="hover:underline">Products</a></li>
      <li><a routerLink="/auth" class="hover:underline">Auth</a></li>
      <li><a routerLink="/orders" class="hover:underline">Orders</a></li>
    </ul>
  </nav>
</header>
<main class="container mx-auto p-4">
  <router-outlet></router-outlet>
</main>
<footer class="bg-gray-800 text-white p-4 text-center">
  <p>&copy; 2025 E-Commerce App</p>
</footer>
```

```typescript
// src/app/app.component.ts (shell-app)
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'shell-app';
}
```

```scss
/* src/app/app.component.scss (shell-app) */
:host {
  display: block;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
main {
  flex: 1;
}
```

```typescript
// src/decl.d.ts (shell-app)
declare module 'productListing/*';
declare module 'userAuth/*';
declare module 'orderHistory/*';
```

```json
// angular.json (shell-app, serve configuration is auto-updated by init)
"The `ng add` command updates the `serve` builder to `"@angular-architects/module-federation:serve"` with port 4200 and publicHost."
```

```javascript
// webpack.config.js (product-listing)
const { shareAll, withModuleFederationPlugin } = require('@angular-architects/module-federation/webpack');

module.exports = withModuleFederationPlugin({
  name: 'productListing',
  exposes: {
    './routes': './src/app/routes.ts'
  },
  shared: {
    ...shareAll({ singleton: true, strictVersion: true, requiredVersion: 'auto' })
  }
});
```

```typescript
// src/app/routes.ts (product-listing)
import { Routes } from '@angular/router';
import { ProductListComponent } from './product-list/product-list.component';

export const ProductListingRoutes: Routes = [
  { path: '', component: ProductListComponent }
];
```

```typescript
// src/app/product-list/product-list.component.ts (product-listing)
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

interface Product {
  id: number;
  name: string;
  price: number;
  image: string;
}

@Component({
  selector: 'app-product-list',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './product-list.component.html',
  styleUrls: ['./product-list.component.scss']
})
export class ProductListComponent {
  products: Product[] = [
    { id: 1, name: 'Laptop', price: 999.99, image: 'laptop.jpg' },
    { id: 2, name: 'Smartphone', price: 499.99, image: 'smartphone.jpg' },
    // Add more mock products
  ];
  filteredProducts: Product[] = [...this.products];
  searchTerm = '';
  page = 1;
  pageSize = 5;

  search() {
    this.filteredProducts = this.products.filter(p =>
      p.name.toLowerCase().includes(this.searchTerm.toLowerCase())
    );
    this.page = 1;
  }

  get paginatedProducts(): Product[] {
    const start = (this.page - 1) * this.pageSize;
    return this.filteredProducts.slice(start, start + this.pageSize);
  }

  get totalPages(): number {
    return Math.ceil(this.filteredProducts.length / this.pageSize);
  }

  changePage(newPage: number) {
    if (newPage >= 1 && newPage <= this.totalPages) {
      this.page = newPage;
    }
  }
}
```

```html
<!-- src/app/product-list/product-list.component.html (product-listing) -->
<div class="p-4">
  <h2 class="text-2xl font-bold mb-4">Products</h2>
  <input
    type="text"
    [(ngModel)]="searchTerm"
    (input)="search()"
    placeholder="Search products..."
    class="border p-2 mb-4 w-full"
  />
  <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <div *ngFor="let product of paginatedProducts" class="border p-4">
      <img [src]="product.image" alt="{{product.name}}" class="h-40 w-full object-cover mb-2" />
      <h3 class="text-lg font-semibold">{{product.name}}</h3>
      <p>${{product.price | number:'1.2-2'}}</p>
    </div>
  </div>
  <div class="flex justify-center mt-4">
    <button (click)="changePage(page - 1)" [disabled]="page === 1" class="px-4 py-2 bg-blue-600 text-white mr-2">Previous</button>
    <span>Page {{page}} of {{totalPages}}</span>
    <button (click)="changePage(page + 1)" [disabled]="page === totalPages" class="px-4 py-2 bg-blue-600 text-white ml-2">Next</button>
  </div>
</div>
```

```scss
/* src/app/product-list/product-list.component.scss (product-listing) */
img {
  max-width: 100%;
}
button:disabled {
  @apply bg-gray-400 cursor-not-allowed;
}
```

```typescript
// src/main.ts (product-listing)
import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';

bootstrapApplication(AppComponent).catch(err => console.error(err));
```

```json
// angular.json (product-listing, serve configuration is auto-updated by init)
"The `ng add` command updates the `serve` builder to `"@angular-architects/module-federation:serve"` with port 4201 and publicHost."
```

```javascript
// webpack.config.js (user-auth)
const { shareAll, withModuleFederationPlugin } = require('@angular-architects/module-federation/webpack');

module.exports = withModuleFederationPlugin({
  name: 'userAuth',
  exposes: {
    './routes': './src/app/routes.ts'
  },
  shared: {
    ...shareAll({ singleton: true, strictVersion: true, requiredVersion: 'auto' })
  }
});
```

```typescript
// src/app/routes.ts (user-auth)
import { Routes } from '@angular/router';
import { AuthComponent } from './auth/auth.component';

export const UserAuthRoutes: Routes = [
  { path: '', component: AuthComponent }
];
```

```typescript
// src/app/auth/auth.component.ts (user-auth)
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-auth',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.scss']
})
export class AuthComponent {
  isLogin = true;
  loginForm = { email: '', password: '' };
  registerForm = { email: '', password: '', confirmPassword: '' };
  errorMessage = '';

  toggleMode() {
    this.isLogin = !this.isLogin;
    this.errorMessage = '';
  }

  login() {
    if (!this.loginForm.email || !this.loginForm.password) {
      this.errorMessage = 'Please fill in all fields';
      return;
    }
    // Mock login logic
    console.log('Login:', this.loginForm);
    this.errorMessage = 'Login successful (mock)';
  }

  register() {
    if (!this.registerForm.email || !this.registerForm.password || !this.registerForm.confirmPassword) {
      this.errorMessage = 'Please fill in all fields';
      return;
    }
    if (this.registerForm.password !== this.registerForm.confirmPassword) {
      this.errorMessage = 'Passwords do not match';
      return;
    }
    // Mock register logic
    console.log('Register:', this.registerForm);
    this.errorMessage = 'Registration successful (mock)';
  }
}
```

```html
<!-- src/app/auth/auth.component.html (user-auth) -->
<div class="p-4 max-w-md mx-auto">
  <h2 class="text-2xl font-bold mb-4">{{isLogin ? 'Login' : 'Register'}}</h2>
  <div *ngIf="errorMessage" class="text-red-600 mb-4">{{errorMessage}}</div>
  
  <div *ngIf="isLogin; else registerForm">
    <form (ngSubmit)="login()">
      <div class="mb-4">
        <label class="block mb-1">Email</label>
        <input type="email" [(ngModel)]="loginForm.email" name="email" class="border p-2 w-full" required />
      </div>
      <div class="mb-4">
        <label class="block mb-1">Password</label>
        <input type="password" [(ngModel)]="loginForm.password" name="password" class="border p-2 w-full" required />
      </div>
      <button type="submit" class="px-4 py-2 bg-blue-600 text-white">Login</button>
    </form>
  </div>
  
  <ng-template #registerForm>
    <form (ngSubmit)="register()">
      <div class="mb-4">
        <label class="block mb-1">Email</label>
        <input type="email" [(ngModel)]="registerForm.email" name="email" class="border p-2 w-full" required />
      </div>
      <div class="mb-4">
        <label class="block mb-1">Password</label>
        <input type="password" [(ngModel)]="registerForm.password" name="password" class="border p-2 w-full" required />
      </div>
      <div class="mb-4">
        <label class="block mb-1">Confirm Password</label>
        <input type="password" [(ngModel)]="registerForm.confirmPassword" name="confirmPassword" class="border p-2 w-full" required />
      </div>
      <button type="submit" class="px-4 py-2 bg-blue-600 text-white">Register</button>
    </form>
  </ng-template>
  
  <button (click)="toggleMode()" class="mt-4 text-blue-600">
    {{isLogin ? 'Switch to Register' : 'Switch to Login'}}
  </button>
</div>
```

```scss
/* src/app/auth/auth.component.scss (user-auth) */
form {
  @apply flex flex-col;
}
button {
  @apply rounded;
}
```

```typescript
// src/main.ts (user-auth)
import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';

bootstrapApplication(AppComponent).catch(err => console.error(err));
```

```json
// angular.json (user-auth, serve configuration is auto-updated by init)
"The `ng add` command updates the `serve` builder to `"@angular-architects/module-federation:serve"` with port 4202 and publicHost."
```

```javascript
// webpack.config.js (order-history)
const { shareAll, withModuleFederationPlugin } = require('@angular-architects/module-federation/webpack');

module.exports = withModuleFederationPlugin({
  name: 'orderHistory',
  exposes: {
    './routes': './src/app/routes.ts'
  },
  shared: {
    ...shareAll({ singleton: true, strictVersion: true, requiredVersion: 'auto' })
  }
});
```

```typescript
// src/app/routes.ts (order-history)
import { Routes } from '@angular/router';
import { OrderListComponent } from './order-list/order-list.component';

export const OrderHistoryRoutes: Routes = [
  { path: '', component: OrderListComponent }
];
```

```typescript
// src/app/order-list/order-list.component.ts (order-history)
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

interface Order {
  id: number;
  date: string;
  amount: number;
}

@Component({
  selector: 'app-order-list',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './order-list.component.html',
  styleUrls: ['./order-list.component.scss']
})
export class OrderListComponent {
  orders: Order[] = [
    { id: 1, date: '2025-10-01', amount: 199.99 },
    { id: 2, date: '2025-09-15', amount: 49.99 },
    // Add more mock orders
  ];
}
```

```html
<!-- src/app/order-list/order-list.component.html (order-history) -->
<div class="p-4">
  <h2 class="text-2xl font-bold mb-4">Order History</h2>
  <table class="w-full border">
    <thead>
      <tr class="bg-gray-100">
        <th class="p-2 text-left">Order ID</th>
        <th class="p-2 text-left">Date</th>
        <th class="p-2 text-left">Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr *ngFor="let order of orders" class="border-t">
        <td class="p-2">{{order.id}}</td>
        <td class="p-2">{{order.date}}</td>
        <td class="p-2">${{order.amount | number:'1.2-2'}}</td>
      </tr>
    </tbody>
  </table>
</div>
```

```scss
/* src/app/order-list/order-list.component.scss (order-history) */
table {
  @apply border-collapse;
}
th, td {
  @apply border;
}
```

```typescript
// src/main.ts (order-history)
import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';

bootstrapApplication(AppComponent).catch(err => console.error(err));
```

```json
// angular.json (order-history, serve configuration is auto-updated by init)
"The `ng add` command updates the `serve` builder to `"@angular-architects/module-federation:serve"` with port 4203 and publicHost."
```

```bash
# Running the Application
# Run each app in separate terminals
cd shell-app
ng serve

cd ../product-listing
ng serve

cd ../user-auth
ng serve

cd ../order-history
ng serve
```

```
# Notes
- The shell app runs on `http://localhost:4200` and lazily loads micro frontend routes.
- All components are now standalone, leveraging Angular 17's architecture for lighter bundles.
- Product Listing app (`http://localhost:4201`) includes search and pagination with template-driven forms.
- User Auth app (`http://localhost:4202`) provides login/register forms with mock validation.
- Order History app (`http://localhost:4203`) displays a table of mock orders.
- Tailwind CSS classes are used for styling where applicable (install via `npm install -D tailwindcss postcss autoprefixer` and configure if needed).
- The apps use Angular 17 and the updated `@angular-architects/module-federation` for integration with standalone components.
- Image sources in Product Listing are placeholders; replace with actual image URLs in a real app.
- Remove any generated `app.module.ts` or `app-routing.module.ts` files as they are not used in standalone setup.
```