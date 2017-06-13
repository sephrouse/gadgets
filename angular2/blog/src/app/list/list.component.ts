import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { FirebaseListObservable } from 'angularfire2';

import { BlogsService } from '../blogs.service';
import { Blog } from '../blog.model';


@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

  blogs: Blog[];

  //constructor(private _blogService: BlogsService) { }
  constructor(private _blogsService: BlogsService, private _router: Router) {}

  ngOnInit() {
    // this._blogService.getBlogs()
    //     .then(response => this.blogs = response);

    //this.blogs = this._blogsService.getBlogs();

    this._blogsService.getBlogs().subscribe(resp => {
      this.blogs = resp;
    });
  }

  readMore(b: Blog) {
    //b.subscribe(b => {this._blogsService.setBlog(b as Blog)})
    //this._blogsService.setBlog(b);
    //this._router.navigate(['/content']);
  }

}
