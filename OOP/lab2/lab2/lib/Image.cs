using System;
using System.Collections.Generic;
using System.Linq;
using Avalonia.Controls;
using System.IO;
using System.Text.Json;
using Avalonia;
using Avalonia.Controls.Shapes;
using Avalonia.Media;
using Avalonia.VisualTree;
using DynamicData;
using lab2.Views;

namespace lab2.lib;

public class Image
{
    public double X { get; set; } = 0;
    public double Y { get; set; } = 0;

    private int _width;
    private int _height;
    
    public int Width
    {
        get => (int) (this._width * this._scale);
        set
        {
            this._width = value;
            if (this.Canvas != null) this.Canvas.Width = this.Width;
        }
    }
    public int Height
    {
        get => (int) (this._height * this._scale);
        set
        {
            this._height = value;
            if (this.Canvas != null) this.Canvas.Height = this.Height;
        }
    }

    public List<Figure> Figures { get; set; }
    public Canvas Canvas;

    private double _scale = 1;
    public double Scale
    {
        get => this._scale;
        set
        {
            this._scale = value;
            this.Width = this._width;
            this.Height = this._height;
        }
    }
    public Figure this[int index]
    {
        get => this.Figures[index];
        set => this.Figures[index] = value;
    }
    
    public string Color { get; set; } = "#AAFF0000";
    
    public Image(double x, double y, int width, int height, string? color = null, List<Figure>? figures = null)
    {
        this.X = x;
        this.Y = y;
        this.Width = width;
        this.Height = height;
        this.Figures = figures ?? [];
        this.Color = color ?? this.Color;
        this.Canvas = new Canvas
        {
            Width = this.Width,
            Height = this.Height,
            ClipToBounds = true,
            Background = Brushes.Transparent,
            Tag = this,
        };
        
        this.Canvas.PointerPressed += MainWindow.SetSelectedImage;
    }

    public Image(string header, List<Figure> figures)
    {
        string[] metadata = header.Split(";");
        this.X = double.Parse(metadata[0]);
        this.Y = double.Parse(metadata[1]);
        this.Width = int.Parse(metadata[2]);
        this.Height = int.Parse(metadata[3]);
        this.Color = metadata[4];
        this.Figures = [];
        this.Canvas = new Canvas
        {
            Width = this.Width,
            Height = this.Height,
            ClipToBounds = true,
            Background = Brushes.Transparent,
            Tag = this,
        };
        
        this.Canvas.PointerPressed += MainWindow.SetSelectedImage;
        
        foreach (var fig in figures)
        {
            this.Add(fig);
        }

        this.Scale = double.Parse(metadata[5]);
    }
    
    public void Move(double x, double y) {
        this.X += x;
        this.Y += y;
        this.DrawAll();
    }
    public void MoveTo(double x, double y) {
        this.X = x;
        this.Y = y;
        this.DrawAll();
    }
    
    public void MoveChildren(double x, double y)
    {
        foreach (var fig in this.Figures)
        {
            fig.Move(x, y);
        }
    }
    public void MoveChildrenTo(double x, double y)
    {
        foreach (var fig in this.Figures)
        {
            fig.MoveTo(x, y);
        }
    }
    
    public void Add(Figure fig)
    {
        fig.ParentImage = this;
        this.Figures.Add(fig);
        this.DrawAll();
    }
    public void RemoveAt(int index)
    {
        this.Figures.RemoveAt(index);
    }

    public void Concat(Image other)
    {
        foreach (var fig in other.Figures)
        {
            this.Add(fig);
        }
    }
    
    public double GetArea()
    {
        return this.Figures.Sum(fig => fig.GetArea());
    }
    public double GetPerimeter()
    {
        return this.Figures.Sum(fig => fig.GetPerimeter());
    }

    public void DrawAll()
    {
        this.Canvas.Children.Clear();
        Avalonia.Controls.Canvas.SetLeft(this.Canvas, this.X);
        Avalonia.Controls.Canvas.SetTop(this.Canvas, this.Y);
        
        this.Canvas.Children.Add(new Line
        {
            StartPoint = new Point(2, 2),
            EndPoint = new Point(2, this.Height - 2),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        this.Canvas.Children.Add(new Line
        {
            StartPoint = new Point(2, 2),
            EndPoint = new Point(this.Width - 2, 2),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        this.Canvas.Children.Add(new Line
        {
            StartPoint = new Point(2, this.Height - 2),
            EndPoint = new Point(this.Width - 2, this.Height - 2),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        this.Canvas.Children.Add(new Line
        {
            StartPoint = new Point(this.Width - 2, 2),
            EndPoint = new Point(this.Width - 2, this.Height - 2),
            Stroke = new SolidColorBrush(Avalonia.Media.Color.Parse(this.Color)),
            StrokeThickness = 2,
        });
        
        foreach (var fig in this.Figures)
        {
            fig.Draw();
        }
    }
    
    public override string ToString()
    {
        return this.Figures.Aggregate(
            $"{this.X};{this.Y};{this.Width};{this.Height};{this.Color};{this.Scale}\n", 
            (current, fig) => current + $"{fig.ToString()}\n"
                );
        
    }
    
    public void Destroy()
    {
        (this.Canvas.GetVisualRoot() as MainWindow)?.Main.Children.Remove(this.Canvas);
    }
}