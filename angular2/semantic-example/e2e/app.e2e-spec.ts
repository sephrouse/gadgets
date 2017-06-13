import { SemanticExamplePage } from './app.po';

describe('semantic-example App', function() {
  let page: SemanticExamplePage;

  beforeEach(() => {
    page = new SemanticExamplePage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
