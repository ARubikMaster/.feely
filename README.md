# .feely image file format



## Linux


* 1: **Download bin/feely**


* 2: **Make the feely executable runnable**
```chmod +x feely```


* 3: **Add executable to path (optional, but reccomended)**

How to do for bash (If you use a different shell look up how to add a folder to path):

```echo 'export PATH="$PATH:/path/to/executable/folder"' >> ~/.bashrc ```


* 4: **To convert a .png image to .feely**

If added to path:

```feely convert /path/to/image.png name_of_new_image```

If not, cd to the executable location and run:

```./feely convert /path/to/image.png name_of_new_image```


## Windows


* 1: **Download bin/feely.exe**


* 2: **Add executable folder to path (optional but recommended)**

You can see an example here: https://windowsloop.com/how-to-add-to-windows-path/

* 3: **To convert a .png image to .feely**


If added to path:

```feely convert C:\path\to\image.png name_of_new_image```

If not, cd to the executable location and run:

```.\feely convert C:\path\to\image.png name_of_new_image```

Please note you might get an antivirus warning, as the executable is not signed.
