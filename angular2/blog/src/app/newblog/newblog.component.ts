import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';

import { BlogsService } from '../blogs.service';
import { Blog } from '../blog.model';

@Component({
  selector: 'app-newblog',
  templateUrl: './newblog.component.html',
  styleUrls: ['./newblog.component.css']
})
export class NewblogComponent implements OnInit {
  verified: boolean = false;

  constructor(private _blogsService: BlogsService, private _router: Router) { }

  ngOnInit() {
  }

  onSubmit(f: NgForm) {
    //this._blogsService.postBlog(f.value);
    //console.log(f.controls['title'].value);
    //console.log(f.controls['content'].value);
    var blog: Blog = {
      title: f.controls['title'].value,
      content: (f.controls['content'].value as string).split('\n'),
      createdDate: Date.now(),
    }

    this._blogsService.postBlog(blog);

    this._router.navigate(['/list']);
  }

  verify(pw: string) {
    console.log(pw);
    
    if (pw === 'q1w2e3r4') {
      this.verified = true;
    }
  }

}
