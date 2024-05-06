using System;
using Avalonia.Controls;
using Avalonia.Media;
using Avalonia.VisualTree;
using lab2.Views;

namespace lab2.lib;

public class Circuit : Figure
{
    private double _radius;
    public double Radius
    {
        get => this.Scaled(this._radius);
        set => this._radius = value;
    }
    
    public Circuit(double radius, double x, double y, string? color = null)
    {
        this.Radius = radius;
        this.X = x;
        this.Y = y;
        this.Color = color ?? this.Color;
    }
    public Circuit(string[] metadata) 
    {
        this.X = double.Parse(metadata[1]);
        this.Y = double.Parse(metadata[2]);
        this.Color = metadata[3];
        this.Scale = double.Parse(metadata[4]);
        this.Radius = double.Parse(metadata[5]);
    }
    
    public override double GetArea() {
        return Math.PI * this.Radius * this.Radius;
    }
    public override double GetPerimeter()
    {
        return 2 * Math.PI * this.Radius;
    }

    public override void Draw()
    {
        var canvas = new Canvas
        {
            Tag = this,
            Width = this.Radius * 2,
            Height = this.Radius * 2,
            Background = Brushes.Transparent,
        };

        canvas.Children.Add(new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.Radius * 2,
            Height = this.Radius * 2,
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            Fill = Brushes.Transparent,
        });
        
        canvas.PointerPressed += MainWindow.SetSelectedFigure;
        
        Canvas.SetLeft(canvas, this.X);
        Canvas.SetTop(canvas, this.Y);
        
        this.ParentImage?.Canvas.Children.Add(canvas);
    }

    public override string ToString()
    {
        return base.ToString() + $";{this.Radius}";
    }

}