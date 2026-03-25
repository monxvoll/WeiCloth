import { NgModule } from '@angular/core'; 
import { RouterModule, Routes } from '@angular/router';

import { Register } from './register/register';
import { Login } from './login/login';

 export const routes: Routes = [
  { path: '', component: Login },
  { path: 'register', component: Register}
 ];

 @NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule] 
})

export class AuthRoutingModule { }