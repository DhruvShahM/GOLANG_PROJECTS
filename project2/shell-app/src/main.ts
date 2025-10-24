import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';
import { provideRouter } from '@angular/router';
import { provideHttpClient } from '@angular/common/http';
import { importProvidersFrom } from '@angular/core';
import { AppRoutingModule } from './app/app-routing.module';

import { initFederation } from '@angular-architects/native-federation';

initFederation('/assets/federation.manifest.json')
  .catch(err => console.error(err))
  .then(async () => {
    bootstrapApplication(AppComponent, {
      providers: [
        provideRouter([]),  // Routes handled by AppRoutingModule
        provideHttpClient(),
        importProvidersFrom(AppRoutingModule),
      ],
    });
  })
  .catch(err => console.error(err));