# copyfolder
A very simple tool, which copies the content of a folder into another folder.


## How to use

copyfolder is very easy to use. Just copy the small programm inside a folder and start the app (in Windows it would be the copyfolder.exe). If there is no configuration file at the same folder a default / example conf file is stored inside the same folder with the name copyfolder.yaml

The file looks like:

```
folders:
- src: path/to/src
  dst: path/to/dst
- src: another/src
  dst: another/dst
```

Just remove the example paths and add your paths. You can provide multiple copy folders.