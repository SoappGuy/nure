using System;
using Avalonia.Controls;
using Avalonia.Media;
using DynamicData;
using lab2.Views;

namespace lab2.lib;

public class Ellipse : Figure
{
    private double _width;
    public double Width
    {
        get => this.Scaled(this._width);
        set => this._width = value;
    }
    
    private double _height;
    public double Height
    {
        get => this.Scaled(this._height);
        set => this._height = value;
    }
    
    public Ellipse(double width, double height, double x, double y, string? color = null)
    {
        this.Width = width;
        this.Height = height;
        this.X = x;
        this.Y = y;
        this.Color = color ?? this.Color;
    }

    public Ellipse(string[] metadata)
    {
        this.X = double.Parse(metadata[1]);
        this.Y = double.Parse(metadata[2]);
        this.Color = metadata[3];
        this.Scale = double.Parse(metadata[4]);
        this.Width = double.Parse(metadata[5]);
        this.Height = double.Parse(metadata[6]);
    }
    
    public override double GetArea() {
        return Math.PI * this.Width/2 * this.Height/2;
    }
    public override double GetPerimeter() {
        return Math.PI * Math.Sqrt((this.Width * this.Width) + (this.Height * this.Height));
    }

    public override void Draw()
    {
        var canvas = new Canvas
        {
            Tag = this,
            // Width = this.Width,
            // Height = this.Height,
            Background = Brushes.Transparent,
        };
        
        canvas.Children.Add(new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.Width,
            Height = this.Height,
            Fill = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
        });
        
        canvas.PointerPressed += MainWindow.SetSelectedFigure;
        
        Canvas.SetLeft(canvas, this.X);
        Canvas.SetTop(canvas, this.Y);
        
        this.ParentImage?.Canvas.Children.Add(canvas);
    }

    public override string ToString()
    {
        return base.ToString() + $";{this.Width};{this.Height}";
    }
}