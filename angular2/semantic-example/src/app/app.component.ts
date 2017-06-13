import { Component, AfterViewInit } from '@angular/core';

declare var $:any;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements AfterViewInit {
  ngAfterViewInit(): void {
    $("div.dropdown").dropdown();
  }

  showModal() {
    $('.ui.basic.modal').modal('show');
  }

      toggleNavigation(): void {
        $("body").toggleClass("mini-navbar");
    }
}
