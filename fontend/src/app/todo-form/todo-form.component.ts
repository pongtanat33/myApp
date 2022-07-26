import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { TodoService } from '../todo.service';
import { Location } from '@angular/common'

@Component({
  selector: 'app-todo-form',
  templateUrl: './todo-form.component.html',
  styleUrls: ['./todo-form.component.scss']
})
export class TodoFormComponent implements OnInit {

 todoFormGroup = new FormGroup({
    id: new FormControl(),
    description: new FormControl(),
    complete: new FormControl(),
    useremail: new FormControl()
  })

  
  
  constructor(    
    private route: ActivatedRoute,
    private todoService:TodoService,
    private location : Location
    ){ }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    
    console.log(id)
    if(id){
      if(this.validateEmail(id!)){
        this.todoFormGroup.patchValue({
          useremail: id
        })
      }
    }
    // if(id)this.todoService.getTodo(id).subscribe(todo=>this.todoFormGroup.setValue(todo))
  }

  onBack(){
    this.location.back()
  }

  onSave(){
    if(this.todoFormGroup.get('id')?.value){
      console.log("up")
      this.todoService.updateTodo(this.todoFormGroup.value).subscribe(()=> this.onBack());
    } else {
      this.todoFormGroup.patchValue({
        complete: false
      })
      console.log("down")
      this.todoService.addNewTodo(this.todoFormGroup.value).subscribe(()=> this.onBack());
    }
  }

  validateEmail(email:string){
    return /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()\.,;\s@\"]+\.{0,1})+([^<>()\.,;:\s@\"]{2,}|[\d\.]+))$/.test(email);
  }

}
