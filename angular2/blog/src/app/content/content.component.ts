import { Component, OnInit, OnDestroy } from '@angular/core';
import { Location } from '@angular/common';
import { ActivatedRoute } from '@angular/router';

import { FirebaseObjectObservable } from 'angularfire2';

import { BlogsService } from '../blogs.service';
import { Blog } from '../blog.model';

@Component({
  selector: 'app-content',
  templateUrl: './content.component.html',
  styleUrls: ['./content.component.css']
})
export class ContentComponent implements OnInit {
  blog: Blog;

  constructor(private _blogsService: BlogsService, private _location: Location, private _router: ActivatedRoute) { }

  ngOnInit() {
    this._router.params.subscribe(key => {
      console.log(key['key']);
      this._blogsService.getBlog(key['key']).subscribe(resp => this.blog = resp);
    });
  }

  goBack() {
    this._location.back();
  }

}
