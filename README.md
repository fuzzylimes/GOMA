![GOMA](web/public/img/Goma-Logo.png "GOMA Logo")

GOMA is a starter project for golang web services, styled with [BULMA css](https://bulma.io/). It provides basic template scaffolding with a build pipeline to support changes to the default BULMA scss.

This project is not associated with BULMA in any way.

## Usage
To use, simply clone the repo, nuke the `.git` folder, and re-init for your project

### Adding Pages
The project comes with a page generator tool to quickly add additional pages and their corresponding routes and handlers.

To use, simply run the `make add_page <page name>` command. No need to include the .html, just the name of the page:

```
make add_page recipes
```

### Modifying Styles
All style changes should be made in the `web/sass/mainstyle.scss` file. 

After making changes, the css can be compiled by running `make build_css`. Generated css will appear in the `web/public/css` directory.

You can learn more about customization with BULMA [in their documentation](https://bulma.io/documentation/customize/concepts/).
