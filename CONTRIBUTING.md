# Contributing

## Adding new icons

_NOTE: This process applies to icons whose acceptance into the repository has already been discussed and approved._

There is a short process to make adding new icons seamless:

1. Clone the project (fork instead if you don't have permission):
    `git clone git@github.com:aquasecurity/cloudsec-icons.git && cd cloudsec-icons`
2. Create a working branch:
    `git checkout -b add-my-awesome-icon`
3. Prepare your icon. The icon should be an SVG file whose name is meaningful and ends with `_Aqua.svg`. The foreground colour should be `#0A00D8`. If you have an SVG in an alternative colour, you should replace all instances of the incorrect foreground colour(s) with this one, perhaps using `s/#[0-9a-f]{6}/#0A00D8/gi`.
4. Add the icon to the `src` directory.
5. Run `make` in the project directory to automatically add the new icon(s) to the table in the README.
6. Add, commit and push your changes:
    `git add src && git commit -am "feat: Add my awesome icon" && git push origin add-my-awesome-icon`
7. Create a pull request on GitHub.
8. Wait for the pull request to be merged.
