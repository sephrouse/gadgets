import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { ListComponent } from './list/list.component';
import { ContentComponent } from './content/content.component';
import { ContactComponent } from './contact/contact.component';
import { AboutComponent } from './about/about.component';
import { NewblogComponent } from './newblog/newblog.component';

@NgModule({
    imports: [
        RouterModule.forRoot([
            { path: 'list', component: ListComponent},
            { path: 'content/:key', component:  ContentComponent},
            { path: 'newblog', component: NewblogComponent},
            { path: 'contact', component: ContactComponent},
            { path: 'about', component: AboutComponent},
            { path: '', redirectTo: '/list', pathMatch: 'full'},
            { path: '**', component: ListComponent}
        ], {useHash: true})
    ],
    exports: [
        RouterModule
    ]
})
export class AppRoutingModule {

}