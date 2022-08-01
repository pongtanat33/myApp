import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { delay, Observable, of, throwError } from 'rxjs';
import { map,tap, catchError} from 'rxjs/operators';
import { Todo } from './todos';
import { environment as env }   from 'src/environments/environment';
// import { environment as prodENV}  from 'src/environments/environment.prod';

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  baseUrl = ""
  constructor(private httpClient: HttpClient) { 
     //this.baseUrl = env.apiUrl_todo
    this.baseUrl = `http://${window.location.hostname}/api/todos`
  }

  getTodos(email:string): Observable<Todo[]> {
    return this.httpClient.get<Todo[]>(this.baseUrl+'/'+email).pipe(
      map((data) => {
        console.log("get all todo",data)
         //You can perform some transformation here
         return data;
      }),
      catchError((err, caught) => {
        console.error(err);
        throw err;
      }
      )
    )
  }

  getTodo(id:number) {
    return this.httpClient.get<Todo[]>(this.baseUrl+'/s/'+id).pipe(
      map((data) => {
        console.log("get single todo",data)
         //You can perform some transformation here
         return data;
      }),
      catchError((err, caught) => {
        console.error(err);
        throw err;
      }
      )
    )
  }

  updateTodo(todo:Todo): Observable<Todo>{
    return this.httpClient.put<Todo>(this.baseUrl,todo)
  }

  addNewTodo(todo:Todo): Observable<Todo>{
    return this.httpClient.post<Todo>(this.baseUrl,todo.description+"-OK")
  }

  deleteTodo(id:number): Observable<void>{
    console.log('delete id',id)
    return this.httpClient.delete<void>(this.baseUrl+'/d/'+id).pipe(
      map((data) => {
        console.log("delete",data)
         //You can perform some transformation here
         return data;
      }),
      catchError((err, caught) => {
        console.error(err);
        throw err;
      }
      )
    )
  }
}
