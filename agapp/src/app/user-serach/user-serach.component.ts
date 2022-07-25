import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '../user.service';
import { Location } from '@angular/common'

@Component({
  selector: 'app-user-serach',
  templateUrl: './user-serach.component.html',
  styleUrls: ['./user-serach.component.scss']
})
export class UserSerachComponent implements OnInit {
  userFormGroup = new FormGroup({
    email: new FormControl(),
    name: new FormControl()
  })

  email = '';
  message = '';
  found = false;
  valid = false;
  constructor(
    public router: Router,
    private userService:UserService,
    private location : Location
    ) { }

  ngOnInit(): void {
    
  }

  onBack(){
    this.valid=false
    this.message=""
  }

  onAddUser(){
    if(this.userFormGroup.get('name')!.value)this.userService.addNewUser(this.userFormGroup.value).subscribe(()=>this.router.navigate(['todos',this.email]));
  }

  onLogin(){
    if(this.validateEmail(this.userFormGroup.get('email')?.value)){
      console.log("enter")
      this.message="";
      this.valid = true;
      console.log(this.userFormGroup.get('email')?.value)
      this.email = this.userFormGroup.get('email')!.value
      this.userService.serachUser(this.email).subscribe({
        next:user=>{this.userFormGroup.setValue(user);this.found=true;this.router.navigate(['todos',this.email]);},
        error: err => {this.message = `This ${this.email} not found please register`;}
      })
    }else{
      this.message= "please enter valid email format"
    }
    // console.log(this.found)
    // if(this.found){
    //   this.router.navigate(['todos',this.email]);
    // }
  }

  validateEmail(email:string){
    return /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()\.,;\s@\"]+\.{0,1})+([^<>()\.,;:\s@\"]{2,}|[\d\.]+))$/.test(email);
  }

}
