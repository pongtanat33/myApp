import { Component, OnInit,Input } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Todo } from '../todos';
import { TodoService } from '../todo.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.scss']
})
export class TodoListComponent implements OnInit {
  @Input() item = ''

  todoFormGroup = new FormGroup({
    email: new FormControl(),
    ID: new FormControl(),
    description: new FormControl(),
    complete: new FormControl()
  })

  todos: Todo[] = [];
  message = '';
  tempid =0;

  constructor( private route: ActivatedRoute,private todoService: TodoService) { }

  ngOnInit(): void {
    const param = this.route.snapshot.paramMap.get('email');
    console.log('param',param)
    if(param)this.item=param
    this.onQeury()
    const todox: Todo ={
      ID:5,
      email:"",
      description:"",
      complete:true
    }
  }

  onQeury(){
    this.todoService.getTodos(this.item)
    .subscribe({
      next: ts => this.todos = ts,
      error: err => this.message = err.message
    });
    console.log("item",this.item)
  }

  onSetComplete(id:number,complete:boolean){
    this.todoFormGroup.patchValue({
      ID:id,
      complete: !complete
    })
    console.log(this.todoFormGroup.value)
    this.todoService.updateTodo(this.todoFormGroup.value).subscribe()
  }

  onAdd(){
    this.onQeury()
  }

  onDelete(id:number,complete:boolean){
    console.log(id)
    this.todoService.deleteTodo(id).subscribe(
      ()=>this.todos = this.todos.filter(todo=> todo.ID !== id)
    );
  }

}
