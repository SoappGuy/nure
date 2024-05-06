using System;
using Avalonia;
using Avalonia.Controls;
using Avalonia.Controls.Shapes;
using Avalonia.Media;
using lab2.Views;

namespace lab2.lib;

public class TruncCone : Figure
{
    private double _topRadius;
    private double _botRadius;
    private double _height;

    public double Height
    {
        get => this.Scaled(_height);
        set => _height = value;
    }
    public double BotRadius
    {
        get => this.Scaled(_botRadius);
        set => _botRadius = value;
    }
    public double TopRadius
    {
        get => this.Scaled(_topRadius);
        set => _topRadius = value;
    }

    private double MaxRadius => Math.Max(this.BotRadius, this.TopRadius);
    private double MinRadius => Math.Min(this.BotRadius, this.TopRadius);
    
    private double BotElWidth => this.BotRadius * 2;
    private double BotElHeight => this.BotRadius;

    private double TopElWidth => this.TopRadius * 2;
    private double TopElHeight => this.TopRadius;
    
    private double L => (this.MaxRadius - this.MinRadius);
    
    public TruncCone(double botRadius, double topRadius, double height, double x, double y, string? color = null)
    {
        this.BotRadius = botRadius;
        this.TopRadius = topRadius;
        this.Height = height;
        this.X = x;
        this.Y = y;
        this.Color = color ?? this.Color;
    }
    public TruncCone(string[] metadata)
    {
        this.X = double.Parse(metadata[1]);
        this.Y = double.Parse(metadata[2]);
        this.Color = metadata[3];
        this.Scale = double.Parse(metadata[4]);
        this.TopRadius = double.Parse(metadata[5]);
        this.BotRadius = double.Parse(metadata[6]);
        this.Height = double.Parse(metadata[7]);
    }
    
    public override double GetArea()
    {
        double trapArea = (this.TopElWidth + this.BotElWidth) / 2 * this.Height;
        double topSemiEllipseArea = (Math.PI * this.TopElHeight/2 * this.TopElWidth/2) / 2;
        double botSemiEllipseArea = (Math.PI * this.BotElHeight/2 * this.BotElWidth/2) / 2;
        
        return trapArea + topSemiEllipseArea + botSemiEllipseArea;
    }
    public override double GetPerimeter()
    {
        double h = Math.Sqrt(this.Height * this.Height + this.L * this.L);

        double botSemiElPer = Math.PI * Math.Sqrt((this.BotElWidth * this.BotElWidth) + (this.BotElHeight * this.BotElWidth)) / 2;
        double topSemiElPer = Math.PI * Math.Sqrt((this.TopElWidth * this.TopElWidth) + (this.TopElHeight * this.TopElWidth)) / 2;

        return h + h + botSemiElPer + topSemiElPer;
    }
    
    public override void Draw()
    {
        var cone = new Canvas{
            Background = Brushes.Transparent,
            Width = this.MaxRadius * 2,
            Height = this.Height + this.MaxRadius / 2 + this.MinRadius / 2,
            Tag = this,
        };
        
        cone.PointerPressed += MainWindow.SetSelectedFigure;
        
        cone.Children.Add(new Line
        {
            StartPoint = new Point(this.MaxRadius, this.TopElHeight / 2),
            EndPoint = new Point(this.MaxRadius, this.Height + this.TopElHeight / 2),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            StrokeDashArray = new Avalonia.Collections.AvaloniaList<double>() { 1, 1 }
        });
        cone.Children.Add(new Line
        {
            StartPoint = new Point(this.MaxRadius - this.TopRadius, this.TopElHeight / 2),
            EndPoint = new Point(this.MaxRadius - this.BotRadius, this.Height + this.TopElHeight / 2),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        cone.Children.Add(new Line
        {
            StartPoint = new Point( 
                this.TopRadius == this.MaxRadius ? this.MaxRadius * 2 : this.TopElWidth + this.L,
                this.TopElHeight / 2
                ),
            EndPoint = new Point(
                this.BotRadius == this.MaxRadius ? this.MaxRadius * 2 : this.BotRadius * 2 + this.L,
                this.Height + this.TopElHeight / 2
                ),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        
        var ellipse1 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.TopElWidth,
            Height = this.TopElHeight,
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        };
        
        Canvas.SetLeft(ellipse1, this.MaxRadius - this.TopRadius);
        cone.Children.Add(ellipse1);

        var ellipse2 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.BotElWidth,
            Height = this.BotElHeight,
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            ZIndex = -3
        };
        var ellipse3 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.BotElWidth,
            Height = this.BotElHeight * 0.9,
            Fill = new SolidColorBrush(Avalonia.Media.Color.Parse("#AA000000")),
            ZIndex = -2
        };
        var ellipse4 = new Avalonia.Controls.Shapes.Ellipse()
        {
            Width = this.BotElWidth,
            Height = this.BotElHeight,
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
            ZIndex = -1,
            StrokeDashArray = new Avalonia.Collections.AvaloniaList<double>() { 1, 1 }
        };
        
        Canvas.SetTop(ellipse2, this.Height - this.BotElHeight / 2 + this.TopElHeight / 2);
        Canvas.SetLeft(ellipse2, this.MaxRadius - this.BotRadius);
        Canvas.SetTop(ellipse3, this.Height - this.BotElHeight / 2 + this.TopElHeight / 2);
        Canvas.SetLeft(ellipse3, this.MaxRadius - this.BotRadius);
        Canvas.SetTop(ellipse4, this.Height - this.BotElHeight / 2 + this.TopElHeight / 2);
        Canvas.SetLeft(ellipse4, this.MaxRadius - this.BotRadius);

        cone.Children.Add(ellipse2);
        cone.Children.Add(ellipse3);
        cone.Children.Add(ellipse4);
        
        Canvas.SetLeft(cone, this.X);
        Canvas.SetTop(cone, this.Y);
        this.ParentImage?.Canvas.Children.Add(cone);
    }
    
    public override string ToString()
    {
        return base.ToString() + $";{this._topRadius};{this._botRadius};{this._height}";
    }
}