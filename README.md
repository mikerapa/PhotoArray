# Overview 
This Go package can be used to create a composite image from a list of images.

##Basic Example
Code  

```Go   
    import "github.com/mikerapa/PhotoArray"  
    func main() {  
        pab := PhotoArray.NewPhotoArrayBuilder()  
        pab.RowLength =3  
        pab.AddPhotoPaths("./testimages/soccer1.jpg", "./testimages/soccer2.jpg", 
            "./testimages/soccer3.jpg", "./testimages/soccer4.jpg", 
            "./testimages/soccer5.jpg", "./testimages/soccer6.jpg")
        
        pab.GenerateArray("./testoutputimages/mycompositephoto.jpg")  
     } `
 ```

Resulting Image  
![Example output](testoutputimages/outputimage130.jpg)

##Properties  

Name|Description|Default Value  
----|-----------|-------------  
RowLength|number of photos on each row|10
PhotoHeight|height of each photo in the array|100
PhotoWidth|width of each photo in the array|100

## Methods

Name|Description|Input parameters|Output values
----|-----------|----------------|-------------  
ClearPaths|remove all paths from the list|none|none  
Length|returns an int indicating the number of paths currently in the list|none|none
AddPhotoPaths|list of paths|string...|error
GenerateArray|Create the photo array with the specified output file path|outputpath string|error