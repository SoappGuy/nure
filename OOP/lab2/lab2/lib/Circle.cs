using Avalonia.Controls;
using Avalonia.Controls.Shapes;
using Avalonia.Media;
using lab2.Views;

namespace lab2.lib;

public class Circle : Circuit
{
    public Circle(double radius, double x, double y, string? color = null) : base(radius, x, y, color) {}
    public Circle(string[] metadata) : base(metadata) {}
    
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
            Fill = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
        });
        
        canvas.PointerPressed += MainWindow.SetSelectedFigure;
        
        Canvas.SetLeft(canvas, this.X);
        Canvas.SetTop(canvas, this.Y);
        
        this.ParentImage?.Canvas.Children.Add(canvas);
    }
}