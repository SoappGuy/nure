using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using DynamicData;
using lab2.lib;

namespace lab2.lib;

public static class Serializer
{
    public static void Write(List<Image> images, string path, string name)
    {
        Write(images, Path.Combine(path, name));
    }
    
    public static void Write(List<Image> images, string path)
    {
        using (StreamWriter outputFile = new StreamWriter(path, false))
        {
            List<string> strings = [];
            foreach (var image in images)
            {
                strings.Add(image.ToString());
            }
            
            outputFile.Write(string.Join("|\n", strings));
        }
    }

    public static List<Image> Read(string path, string name)
    {
        return Read(Path.Combine(path, name));
    }
    
    public static List<Image> Read(string path)
    {
        List<Image> images = [];
    
        using (StreamReader inputFile = new StreamReader(path))
        {
            foreach (var imageStrings in inputFile.ReadToEnd()[..^1].Split("\n|\n"))
            {
                string[] strings = imageStrings.Split("\n");
                
                List <Figure> figures = [];
                foreach (var figString in strings[1..])
                {
                    string[] metadata = figString.Split(";");
                    switch (metadata[0])
                    {
                        case "Dot":
                            figures.Add(new Dot(metadata));
                            break;
                        case "Circuit":
                            figures.Add(new Circuit(metadata));
                            break;
                        case "Circle":
                            figures.Add(new Circle(metadata));
                            break;
                        case "Ellipse":
                            figures.Add(new Ellipse(metadata));
                            break;
                        case "Cone":
                            figures.Add(new Cone(metadata));
                            break;
                        case "TruncCone":
                            figures.Add(new TruncCone(metadata));
                            break;
                        default:
                            throw new Exception($"Unrecognized figure \"{metadata[0]}\"");
                    }
                }

                images.Add(new Image(strings[0], figures));
            }
        }
        
        return images;
    }
}
