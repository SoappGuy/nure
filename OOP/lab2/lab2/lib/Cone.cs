using System;
using Avalonia;
using Avalonia.Controls;
using Avalonia.Controls.Shapes;
using Avalonia.Layout;
using Avalonia.Media;
using DynamicData;
using lab2.Views;

namespace lab2.lib;

public class Cone : Figure
{
    
    private double _height;
    private double _radius;
    
    public double Radius
    {
        get => this.Scaled(_radius);
        set => _radius = value;
    }
    public double Height
    {
        get => this.Scaled(_height);
        set => _height = value;
    }
    
    private double ElWidth => this.Radius * 2;
    private double ElHeight => this.Radius;

    public Cone(double radius, double height, double x, double y, string? color = null)
    {
        this.Radius = radius;
        this.Height = height;
        this.X = x;
        this.Y = y;
        this.Color = color ?? this.Color;
    }

    public Cone(string[] metadata)
    {
        this.X = double.Parse(metadata[1]);
        this.Y = double.Parse(metadata[2]);
        this.Color = metadata[3];
        this.Scale = double.Parse(metadata[4]);
        this.Radius = double.Parse(metadata[5]);
        this.Height = double.Parse(metadata[6]);
    }
    
    public override double GetArea()
    {
        double l = Math.Sqrt(this.Radius * this.Radius + this.Height * this.Height);

        double triangleArea = 2 * this.Height * l;
        double semiEllipseArea = (Math.PI * this.ElHeight/2 * this.ElWidth/2) / 2;
        
        return triangleArea + semiEllipseArea;
    }
    public override double GetPerimeter()
    {
        double l = Math.Sqrt(this.Radius * this.Radius + this.Height * this.Height);
        double semiEllipsePerimeter = Math.PI * Math.Sqrt((this.ElWidth * this.ElWidth) + (this.ElHeight * this.ElHeight)) / 2;
        
        return l + l + semiEllipsePerimeter;
    }

    public override string ToString()
    {
        return base.ToString() + $";{this._radius};{this._height}";
    }

    public override void Draw()
    {
        var cone = new Canvas{
            Background = Brushes.Transparent,
            Tag = this,
            Width = this.Radius * 2,
            Height = this.Height + this.Radius / 2,
        };
        
        cone.PointerPressed += MainWindow.SetSelectedFigure;
        
        cone.Children.Add(new Line
        {
            StartPoint = new Point(this.Radius, 0),
            EndPoint = new Point(this.Radius, this.Height),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            StrokeDashArray = new Avalonia.Collections.AvaloniaList<double>() { 1, 1 }
        });
        
        cone.Children.Add( new Line
        {
            StartPoint = new Point(this.Radius, 0),
            EndPoint = new Point(0, this.Height),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        
        cone.Children.Add( new Line
        {
            StartPoint = new Point(this.Radius, 0),
            EndPoint = new Point((2 * this.Radius), this.Height),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });

        var ellipse1 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.ElWidth,
            Height = this.ElHeight,
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            ZIndex = -3,
        };

        var ellipse2 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.ElWidth,
            Height = this.ElHeight,
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            ZIndex = -1,
            StrokeDashArray = new Avalonia.Collections.AvaloniaList<double>() { 1, 1 }
        };
        
        var ellipse3 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.ElWidth,
            Height = this.ElHeight * 0.9,
            Fill = new SolidColorBrush(Avalonia.Media.Color.Parse("#AA000000")),
            ZIndex = -2
        };
        
        Canvas.SetTop(ellipse1, this.Height - (this.ElHeight / 2));
        Canvas.SetTop(ellipse2, this.Height - (this.ElHeight / 2));
        Canvas.SetTop(ellipse3, this.Height - (this.ElHeight / 2));

        cone.Children.Add(ellipse1);
        cone.Children.Add(ellipse2);
        cone.Children.Add(ellipse3);
        
        Canvas.SetLeft(cone, this.X * (this.ParentImage?.Scale ?? 1));
        Canvas.SetTop(cone, this.Y * (this.ParentImage?.Scale ?? 1));
        this.ParentImage?.Canvas.Children.Add(cone);
    }
}