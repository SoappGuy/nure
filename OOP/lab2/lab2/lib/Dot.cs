using Avalonia.Controls;
using Avalonia.Media;
using lab2.Views;

namespace lab2.lib;

public class Dot : Figure
{
    private double _radius => this.Scaled(2.5);

    public Dot(double x, double y, string? color = null)
    {
        this.X = x;
        this.Y = y;
        this.Color = color ?? this.Color;
    }

    public Dot(string[] metadata)
    {
        this.X = double.Parse(metadata[1]);
        this.Y = double.Parse(metadata[2]);
        this.Color = metadata[3];
        this.Scale = double.Parse(metadata[4]);
    }
    
    public override double GetArea()
    {
        return 0;
    }

    public override double GetPerimeter()
    {
        return 0;
    }

    public override void Draw()
    {
        var canvas = new Canvas
        {
            Tag = this,
            Width = this._radius * 2,
            Height = this._radius * 2,
            Background = Brushes.Transparent,
        };

        canvas.Children.Add(new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this._radius * 2,
            Height = this._radius * 2,
            Fill = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
        });
        
        canvas.PointerPressed += MainWindow.SetSelectedFigure;
        
        Canvas.SetLeft(canvas, this.X);
        Canvas.SetTop(canvas, this.Y);
        
        this.ParentImage?.Canvas.Children.Add(canvas);
    }
}