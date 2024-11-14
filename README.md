# PERSONAL CDN

## Background

Throughout time I have noticed that I use the same static files like CSS. In past times, prior to this project, I used to copy and paste the same static files from one project to another which lead to inconsistencies and higher maintenance. As I continue to develop my own style and design systems, I realized I needed a centralized place for this files to be shared across all my projects both personal and commercial. Therefore, I decided to create this small CDN to serve all my static files.

## Why am I not using a CDN

A CDN would be the way to go, but at the time of writing, this is overkilled. This would require another third party app to learn and another monthly subscription. Neither of those am I willing to undertake for now. In the future, if managing a custom server proves greater work, this might be my next option.

## Why am I not using a centralized GIT repo or an package manager

Both of this options are great for version control but not for accessibility. I want to make changes to the files and have them immediately reflect in the projects that are using them. Rather then opening those project and pulling the latest versions.

## Why am I not using a Symlink since all of my projects exist in the same machine

For the same reason I am not using GIT and a package manager

## Purpose

To serve static files across all my projects, commercial and personal. Only those files that are exactly the same in all projects and do not overwrite the particularity of the projects are served here. The main purpose, however, at the time of writing is to share basic css, js, and html styles across my apps.

### About CSS

This project includes skeleton styles that aid but do not generally intrude other projects. The styles provided by these sheets are basic styles that can be built upon and overwritten for greater customization. There are three types of style sheets provided at this time. Those are:

- icons: A small library of commonly used icons
- tokens: Default dictionary tokens for colors, sizes, fonts, etc.
- utils: A set of basic utility classes that reflect bootstrap naming patterns

To build on top of this system, simply overwrite the dictionary tokens.

## Dev

- Simply run `go run main.go` and all the magic should happen ✨
- all static files to be served can be scp'd from the `src/dist` dir by running the command `send_dist_to_remote.sh`. This is the file served by nginx
- To get all the stats of the VPS run `get_vps_project_stats.sh`. This script will fetch the data and write to the html file in the `/dist` directory
- **OR** run ./deploy.sh to execute the entire process
- **THIS IS NOT A SERVER OR APP** It is only a minifying, bundling, and asset-builder for the files to be served
- ~~The ionic icons are located in their own directory `iconic`. Bundle it and copy the `dist` folder in the `/dist/ionic` dire in this project.~~
- I have removed the ioninc icons. Icons are now served from each app root's public directory. The compiled dist comes from the iconic repo.
- The app.go is a small server in charge of serving public files to the client

### About main.go

This files to be minimized are hardcoded in main.go. If adding or deleting files/ dirs, make sure to edit this file and rebuild the binary with `go build`

## TO DO

- Add some styles to the `src/dist/html/server_stats.html`
- Make a visual page of all the colors, icons and assets used
- ~~Find a better system for coloring svg icons~~
- Add the dist files from ionic icons to this project and serve them from here
