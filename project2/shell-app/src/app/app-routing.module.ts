import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { loadRemoteModule } from '@angular-architects/native-federation';

const routes: Routes = [
  {
    path: 'products',
    loadChildren: () =>
      loadRemoteModule('productListing', './Module').then(m => m.ProductListingModule)
  },
  {
    path: 'auth',
    loadChildren: () =>
      loadRemoteModule('userAuth', './Module').then(m => m.UserAuthModule)
  },
  {
    path: 'orders',
    loadChildren: () =>
      loadRemoteModule('orderHistory', './Module').then(m => m.OrderHistoryModule)
  },
  { path: '', redirectTo: '/products', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {}