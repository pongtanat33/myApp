import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { User } from '../users';
import { UserService } from '../user.service';

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss']
})
export class UserListComponent implements OnInit {

  heroFormGroup = new FormGroup({
    email: new FormControl(),
    name: new FormControl()
  })

  message = '';

  users: User[] = [];

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getUsers()
    .subscribe({
      next: us => this.users = us,
      error: err => this.message = err.message
    });
    console.log("myhero",this.users)
  }

}
