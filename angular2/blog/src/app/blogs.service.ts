import { Injectable } from '@angular/core';

import { AngularFire, FirebaseListObservable, FirebaseObjectObservable } from 'angularfire2';

import { Blog } from './blog.model';

@Injectable()
export class BlogsService {
  blogs: FirebaseListObservable<Blog[]>;
  blog: FirebaseObjectObservable<Blog>;

  constructor(private _angularFire: AngularFire) { 
    this.blogs = _angularFire.database.list('example');
  }

  getBlogs(): FirebaseListObservable<Blog[]> {
    return this.blogs;
  }

  postBlog(blog: Blog) {
    this.blogs.push(blog);
    // need to process the exception when the push is failed.
  }

  getBlog(key: string): FirebaseObjectObservable<Blog> {
    this.blog = this._angularFire.database.object('example/' + key);
    return this.blog;
  }

  // getBlog(key: string) {
  //   console.log('service:', key);
  //   this._angularFire.database.object('example/' + key).subscribe(data => {
  //     console.log(JSON.stringify(data));

  //   });
  // }

}
