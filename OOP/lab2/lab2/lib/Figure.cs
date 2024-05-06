using System;
using System.IO;
using Avalonia.Controls;

namespace lab2.lib;

public abstract class Figure
{
    public double X { get; set; } = 0;
    public double Y { get; set; } = 0;
    public string Color { get; set; } = "#FFFF0000";
    public double Scale { get; set; } = 1;
    public Image? ParentImage;
    
    public void Move(double x = 0, double y = 0) {
        this.X += x;
        this.Y += y;
        this.ParentImage?.DrawAll();
    }
    public void MoveTo(double x, double y) {
        this.X = x;
        this.Y = y;
        this.ParentImage?.DrawAll();
    }
    
    public abstract double GetArea();
    public abstract double GetPerimeter();

    public new virtual string ToString()
    {
        return $"{this.GetType().Name};{this.X};{this.Y};{this.Color};{this.Scale}";
    }

    public abstract void Draw();

    public double Scaled(double property)
    {
        return property * this.Scale * (this.ParentImage?.Scale ?? 1);
    }

    public void Destroy()
    {
        this.ParentImage?.Figures.Remove(this);
        this.ParentImage?.DrawAll();
    }
}