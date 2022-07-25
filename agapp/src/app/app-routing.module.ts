import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule ,Routes } from '@angular/router';
import { UserListComponent } from './user-list/user-list.component';
import { TodoListComponent } from './todo-list/todo-list.component';
import { LoginComponent } from './login/login.component';
import { TodoFormComponent } from './todo-form/todo-form.component';
const routes: Routes = [
  {path:'users',component:UserListComponent},
  {path:'todos',component:TodoListComponent},
  {path:'todos/:email',component:TodoListComponent},
  {path:'login',component:LoginComponent},
  {path:'todo-form',component:TodoFormComponent},
  {path:'todo-form/:id',component:TodoFormComponent},
]

@NgModule({
  declarations: [],
  imports: [
    RouterModule.forRoot(routes)
  ],
  exports:[
    RouterModule
  ]
})
export class AppRoutingModule { }
