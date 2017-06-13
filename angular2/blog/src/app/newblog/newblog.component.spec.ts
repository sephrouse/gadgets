/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { NewblogComponent } from './newblog.component';

describe('NewblogComponent', () => {
  let component: NewblogComponent;
  let fixture: ComponentFixture<NewblogComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewblogComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewblogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
