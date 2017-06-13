import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AngularFireModule } from 'angularfire2';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';
import { MDLUpgradeElementDirective } from './mdl-upgrade-element.directive';
import { ListComponent } from './list/list.component';
import { ContentComponent } from './content/content.component';
import { ContactComponent } from './contact/contact.component';
import { AboutComponent } from './about/about.component';
import { NewblogComponent } from './newblog/newblog.component';

import { BlogsService} from './blogs.service';


export const firebaseConfig = {
  apiKey: "AIzaSyDYZFDxane-lUk91ujOvaGPLJfoVzZ7CME",
  authDomain: "syneyshengithubpage.firebaseapp.com",
  databaseURL: "https://syneyshengithubpage.firebaseio.com",
  storageBucket: "syneyshengithubpage.appspot.com",
  messagingSenderId: "414413949289"
};


@NgModule({
  declarations: [
    AppComponent,
    MDLUpgradeElementDirective,
    ListComponent,
    ContentComponent,
    ContactComponent,
    AboutComponent,
    NewblogComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    AppRoutingModule,
    AngularFireModule.initializeApp(firebaseConfig)
  ],
  providers: [BlogsService],
  bootstrap: [AppComponent]
})
export class AppModule { }
