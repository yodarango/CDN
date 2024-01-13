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

To serve static files across all my project, commercial and personal. Only those files that are exactly the same in all projects and do not overwrite the particularity of the projects are served here.
