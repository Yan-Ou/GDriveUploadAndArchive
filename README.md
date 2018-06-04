# GDriveUploadAndArchive
 

* It uses Google Drive REST API to upload/management the JPG files. https://developers.google.com/drive/api/v2/reference/ 

* Follow this guide to enable Good Drive API and download the client_secret.json file for retreiving OAuth token.

* It uses Go Walk func to traverse the folder and subfolders. https://golang.org/pkg/path/filepath/#Walk 

* It uses Go zip to archive all the files once the uploading to Google Drive is completed. https://golang.org/pkg/archive/zip/ 
