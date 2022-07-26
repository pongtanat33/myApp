import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { delay, Observable, of, throwError } from 'rxjs';
import { map,tap, catchError} from 'rxjs/operators';
import { User } from './users';
import { environment as env }   from 'src/environments/environment';
// import { environment as prodENV}  from 'src/environments/environment.prod';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  
  baseUrl = ""
  constructor(private httpClient: HttpClient) {
    // if(devENV.production){
    //   this.baseUrl = devENV.apiUrl
    // }else{
    //   this.baseUrl = prodENV.apiUrl
    // }
    console.log(env.production)
    // this.baseUrl = env.apiUrl
    this.baseUrl = `http://${window.location.hostname}/api/users`
    console.log(this.baseUrl )
    console.log('$userlocation',window.location.hostname);
   }

   addNewUser(user:User): Observable<User>{
    return this.httpClient.post<User>(this.baseUrl,user)
  }

  getUsers(): Observable<User[]> {
    return this.httpClient.get<User[]>(this.baseUrl).pipe(
      map((data) => {
        console.log("mydata",data)
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

  serachUser(email:string): Observable<User>{
    return this.httpClient.get<User>(this.baseUrl+'/'+email).pipe(
      map((data) => {
        console.log("search data",data)
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
