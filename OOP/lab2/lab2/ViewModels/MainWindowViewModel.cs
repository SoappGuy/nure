using System;
using Avalonia.Collections;
using lab2.Views;

namespace lab2.ViewModels;

public class MainWindowViewModel : ViewModelBase
{
    public bool IsSaved => MainWindow.IsSaved;
    public string IsSavedStr => IsSaved ? "" : "*";
    public string Path => MainWindow.Path;
    public string Title => $"{Path.Split('/')[^1]}{IsSavedStr}";
    public string SelectedItem => MainWindow.SelectedFigure?.GetType().Name ?? MainWindow.SelectedImage?.GetType().Name ?? "";
    
    public AvaloniaList<Type> Types { get; set; } =
    [
        typeof(lib.Circuit), typeof(lib.Circle), typeof(lib.Ellipse), typeof(lib.Cone), typeof(lib.TruncCone),
        typeof(lib.Image)
    ];
}