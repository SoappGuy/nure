using Avalonia.Controls;
using System;
using System.Collections.Generic;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using Avalonia;
using Avalonia.Collections;
using Avalonia.Input;
using Avalonia.Interactivity;
using Avalonia.Media;
using Avalonia.VisualTree;
using DynamicData;
using lab2.lib;

namespace lab2.Views;

public partial class MainWindow : Window
{
    public static lib.Image? SelectedImage;
    public static lib.Figure? SelectedFigure;
    public static bool IsSaved = false;
    public static string? Path = null;
    public static Dictionary<string, TextBox> Inputs = new();
    private Dictionary<Key, bool> _keyStates = new();
    
    public MainWindow()
    {
        InitializeComponent();
        this.KeyDown += OnKeyDown;
        this.KeyUp += OnKeyUp;
    }
    
    private void CreateUiForType(Type objectType, Panel parentPanel)
    {
        EditPanel.Children.Clear();
        SpawnPanel.Children.Clear();
        Inputs.Clear();
        
        parentPanel.Children.Add(new TextBlock
        {
            Text = $"{objectType.Name}:",
            Foreground = Brushes.Bisque,
            Padding = Thickness.Parse("3 0 0 0"),
            FontSize = 16,
        });

        if (SelectedFigure != null)
        {
            parentPanel.Children.Add(new TextBlock
            {
                Text = $"Area: {SelectedFigure.GetArea():F2}\nPerimeter: {SelectedFigure.GetPerimeter():F2}",
                Foreground = Brushes.Bisque,
                Padding = Thickness.Parse("3 0 0 0"),
                FontSize = 10,
            });
        }
        else if (SelectedImage != null)
        {
            parentPanel.Children.Add(new TextBlock
            {
                Text = $"Area: {SelectedImage.GetArea():F2}\nPerimeter: {SelectedImage.GetPerimeter():F2}",
                Foreground = Brushes.Bisque,
                Padding = Thickness.Parse("3 0 0 0"),
                FontSize = 14,
            });
        }

        var properties = objectType.GetProperties();
        foreach (var property in properties)
        {
            var label = new Label { Content = property.Name };
            var textBox = new TextBox { Name = $"txt{property.Name}", Tag = property.Name };

            switch (property.Name)
            {
                case "Color":
                    textBox.LostFocus += HexColor_TextInput;       
                    break;
                case "Scale":
                case "X":
                case "Y":
                case "Width":
                case "Height":
                case "Radius":
                case "BotRadius":
                case "TopRadius":
                    textBox.LostFocus += NumberInput_TextInput;
                    break;
                case "Item":
                case "Figures":
                    continue;
                    break;
                default:
                    break;
            }
            
            if (parentPanel.Name == "EditPanel")
            {
                if (SelectedFigure != null)
                {
                    textBox.Watermark = property.GetValue(SelectedFigure)?.ToString() ?? "";
                }
                else if (SelectedImage != null)
                {
                    textBox.Watermark = property.GetValue(SelectedImage)?.ToString() ?? "";
                }
            }
            
            parentPanel.Children.Add(label);
            parentPanel.Children.Add(textBox);
            Inputs.Add(textBox.Name, textBox);

            if (parentPanel.Name == "SpawnPanel")
            {
                SpawnButton.IsVisible = true;
                
                ApplyButton.IsVisible = false;
                DeleteButton.IsVisible = false;
            } else if (parentPanel.Name == "EditPanel")
            {
                ApplyButton.IsVisible = true;
                DeleteButton.IsVisible = true;
                
                SpawnButton.IsVisible = false;
                SpawnComboBox.SelectedIndex = 0;
            }
            
        }
    }
    
    public void OnObjectTypeSelected(object sender, SelectionChangedEventArgs e)
    {
        ComboBox? parsedSender = (ComboBox)sender;
        string? selectedTypeName = ((ComboBoxItem?)parsedSender.SelectedItem)?.Content?.ToString() ?? "";
        Type? type = null;

        if (SpawnButton != null)
        {
            SpawnButton.IsVisible = true;
        }
        
        switch (selectedTypeName)
        {
            case "Circuit":
                type = typeof(Circuit);
                break;
            case "Circle":
                type = typeof(Circle);
                break;
            case "Ellipse":
                type = typeof(Ellipse);
                break;
            case "Cone":
                type = typeof(Cone);
                break;
            case "TruncCone":
                type = typeof(TruncCone);
                break;
            case "Image":
                type = typeof(lib.Image);
                break;
            case "None":
                if (SpawnButton != null)
                {
                    SpawnButton.IsVisible = false;
                }
                break;
            default:
                break;
        }

        if (type == null)
        {
            return;
        }
        
        CreateUiForType(type, SpawnPanel);
    }
    
    private bool ValidateColor(string colorString)
    {
        return new Regex("^#[0-9A-Fa-f]{0,8}$").IsMatch(colorString);
    }
    
    private bool ValidateNumber(string numString)
    {
        return new Regex(@"^-?\d+(\.\d+)?$").IsMatch(numString);
    }

    private void HexColor_TextInput(object? sender, RoutedEventArgs e)
    {
        var textBox = sender as TextBox;
        if (textBox.Text == null) return;
        
        string text = textBox.Text;
        if (!ValidateColor(text))
        {
            textBox.Text = "";
            ShowError($"Can't parse color from [{text}]!");
        }
    }

    private void NumberInput_TextInput(object? sender, RoutedEventArgs e)
    {
        var textBox = sender as TextBox;
        if (textBox.Text == null) return;

        string text = textBox.Text;
        if (!ValidateNumber(text))
        {
            textBox.Text = "";
            ShowError($"Can't parse number from [{text}]!");
        }
    }

    private void Spawn_OnClick(object? sender, RoutedEventArgs e)
    {
        var selected = SpawnComboBox.SelectedItem as ComboBoxItem;
        Figure? fig = null;
        lib.Image? img = null;
        
        bool isTopRadius = Inputs.TryGetValue("txtTopRadius",  out TextBox? txtTopRadius) && !string.IsNullOrEmpty(txtTopRadius.Text);
        bool isBotRadius = Inputs.TryGetValue("txtBotRadius",  out TextBox? txtBotRadius) && !string.IsNullOrEmpty(txtBotRadius.Text);
        bool isRadius =    Inputs.TryGetValue("txtRadius",     out TextBox? txtRadius)    && !string.IsNullOrEmpty(txtRadius.Text);
        bool isHeight =    Inputs.TryGetValue("txtHeight",     out TextBox? txtHeight)    && !string.IsNullOrEmpty(txtHeight.Text);
        bool isWidth =     Inputs.TryGetValue("txtWidth",      out TextBox? txtWidth)     && !string.IsNullOrEmpty(txtWidth.Text);
        bool isScale =    Inputs.TryGetValue("txtScale",       out TextBox? txtScale)     && !string.IsNullOrEmpty(txtScale.Text);
        bool isColor =     Inputs.TryGetValue("txtColor",      out TextBox? txtColor)     && !string.IsNullOrEmpty(txtColor.Text);
        bool isX =         Inputs.TryGetValue("txtX",          out TextBox? txtX)         && !string.IsNullOrEmpty(txtX.Text);
        bool isY =         Inputs.TryGetValue("txtY",          out TextBox? txtY)         && !string.IsNullOrEmpty(txtY.Text);
        
        double topRadius = double.TryParse(txtTopRadius?.Text, out double _topRadius) ? _topRadius : 25;
        double botRadius = double.TryParse(txtBotRadius?.Text, out double _botRadius) ? _botRadius : 50;
        double radius =    double.TryParse(txtRadius?.Text,    out double _radius)    ? _radius    : 25;
        double height =    double.TryParse(txtHeight?.Text,    out double _height)    ? _height    : 75;
        double width =     double.TryParse(txtWidth?.Text,     out double _width)     ? _width     : 100;
        double scale =     double.TryParse(txtScale?.Text,     out double _scale)     ? _scale     : 1;
        double x =         double.TryParse(txtX?.Text,         out double _x)         ? _x         : 0;
        double y =         double.TryParse(txtY?.Text,         out double _y)         ? _y         : 0;
        string color =     isColor ? txtColor.Text : "#FFFF0000";
    
        switch (selected?.Content)
        {
            case "Circuit":
                fig = new Circuit(radius, x, y, color);
                break;
            case "Circle":
                fig = new Circle(radius, x, y, color);
                break;
            case "Ellipse":
                fig = new Ellipse(width, height, x, y, color);
                break;
            case "Cone":
                fig = new Cone(radius, height, x, y, color);
                break;
            case "TruncCone":
                fig = new TruncCone(botRadius, topRadius, height, x, y, color);
                break;
            case "Image":
                img = new lib.Image(x, y, (int)width, (int)height, color);
                break;
            default:
                ShowError("Unexpected Figure");
                return;
                break;
        }

        if (selected.Content?.ToString() != "Image")
        {
            if (isScale) fig.Scale = scale;
            if (SelectedImage != null)
            {
                SelectedImage.Add(fig);
            }
            else
            {
                img = new lib.Image(0, 0, 400, 400);
                Main.Children.Add(img.Canvas);
                img.Add(fig);
                SelectedImage = img;
            }

            SelectedFigure = fig;
        }
        else
        {
            if (isScale) img.Scale = scale;
            Main.Children.Add(img.Canvas);
            img.DrawAll();
            SelectedImage = img;
            SelectedFigure = null;
        }

        IsSaved = false;
    }

    private void Edit_OnClick(object? sender, RoutedEventArgs e)
    {
        bool isTopRadius = Inputs.TryGetValue("txtTopRadius",  out TextBox? txtTopRadius) && !string.IsNullOrEmpty(txtTopRadius.Text);
        bool isBotRadius = Inputs.TryGetValue("txtBotRadius",  out TextBox? txtBotRadius) && !string.IsNullOrEmpty(txtBotRadius.Text);
        bool isRadius =    Inputs.TryGetValue("txtRadius",     out TextBox? txtRadius)    && !string.IsNullOrEmpty(txtRadius.Text);
        bool isHeight =    Inputs.TryGetValue("txtHeight",     out TextBox? txtHeight)    && !string.IsNullOrEmpty(txtHeight.Text);
        bool isWidth =     Inputs.TryGetValue("txtWidth",      out TextBox? txtWidth)     && !string.IsNullOrEmpty(txtWidth.Text);
        bool isScale =    Inputs.TryGetValue("txtScale",       out TextBox? txtScale)     && !string.IsNullOrEmpty(txtScale.Text);
        bool isColor =     Inputs.TryGetValue("txtColor",      out TextBox? txtColor)     && !string.IsNullOrEmpty(txtColor.Text);
        bool isX =         Inputs.TryGetValue("txtX",          out TextBox? txtX)         && !string.IsNullOrEmpty(txtX.Text);
        bool isY =         Inputs.TryGetValue("txtY",          out TextBox? txtY)         && !string.IsNullOrEmpty(txtY.Text);
        
        double.TryParse(txtTopRadius?.Text, out double topRadius);
        double.TryParse(txtBotRadius?.Text, out double botRadius);
        double.TryParse(txtHeight?.Text,    out double height);
        double.TryParse(txtWidth?.Text,     out double width);
        double.TryParse(txtScale?.Text,     out double scale);
        double.TryParse(txtRadius?.Text,    out double radius);
        double.TryParse(txtX?.Text,         out double x);
        double.TryParse(txtY?.Text,         out double y);
        string color = txtColor?.Text ?? "";
        
        dynamic? fig = null;
        string figType = SelectedFigure?.GetType().Name ?? SelectedImage?.GetType().Name ?? "None";

        if (figType == "None")
        {
            return;
        }
        
        switch (figType)
        {
            case "Circuit":
                fig = SelectedFigure as Circuit;
                fig.Radius = isRadius ? radius : fig.Radius;
                break;
            case "Circle":
                fig = SelectedFigure as Circle;
                fig.Radius = isRadius ? radius : fig.Radius;
                break;
            case "Ellipse":
                fig = SelectedFigure as Ellipse;
                fig.Height = isHeight ? height : fig.Height;
                fig.Width = isWidth ? width : fig.Width;
                break;
            case "Cone":
                fig = SelectedFigure as Cone;
                fig.Radius = isRadius ? radius : fig.Radius;
                fig.Height = isHeight ? height : fig.Height;
                break;
            case "TruncCone":
                fig = SelectedFigure as TruncCone;
                fig.BotRadius = isBotRadius ? botRadius : fig.BotRadius;
                fig.TopRadius = isTopRadius ? topRadius : fig.TopRadius;
                fig.Height = isHeight ? height : fig.Height;
                break;
            case "Image":
                fig = SelectedImage as lib.Image;
                fig.Height = isHeight ? (int)height : fig.Height;
                fig.Width = isWidth ? (int)width : fig.Width;
                break;
            default:
                ShowError("Unexpected Figure");
                break;
        }

        if (fig != null)
        {
            fig.Color = isColor ? color : fig.Color;
            fig.Scale = isScale ? scale : 1;
            fig.X = isX ? x : fig.X;
            fig.Y = isY ? y : fig.Y;
        }
        
        SelectedImage?.DrawAll();
    }
    
    public static void SetSelectedImage(object? sender, PointerPressedEventArgs e)
    {
        SelectedFigure = null;
        SelectedImage = (sender as Canvas)?.Tag as lib.Image ?? null;

        e.Handled = true;
        
        if (SelectedImage != null)
        {
            var mainWindow = SelectedImage.Canvas.GetVisualRoot() as MainWindow;
            mainWindow.CreateUiForType(SelectedImage.GetType(), mainWindow.EditPanel);
        }
    }
    
    public static void SetSelectedFigure(object? sender, PointerPressedEventArgs e)
    {
        var tagged = (sender as Canvas)?.Tag as lib.Figure; 
        SelectedImage = tagged?.ParentImage ?? null;
        SelectedFigure = tagged ?? null;
        
        e.Handled = true;
        
        if (SelectedFigure != null)
        {
            var mainWindow = SelectedImage.Canvas.GetVisualRoot() as MainWindow;
            mainWindow.CreateUiForType(SelectedFigure.GetType(), mainWindow.EditPanel);
        }
    }
    
    private void OnKeyUp(object sender, KeyEventArgs e)
    {
        _keyStates[e.Key] = false;
    }
    
    private void OnKeyDown(object? sender, KeyEventArgs e)
    {
        _keyStates[e.Key] = true;

        _keyStates.TryGetValue(Key.W, out bool W);
        _keyStates.TryGetValue(Key.A, out bool A);
        _keyStates.TryGetValue(Key.S, out bool S);
        _keyStates.TryGetValue(Key.D, out bool D);
        _keyStates.TryGetValue(Key.LeftCtrl, out bool LeftCtrl);
        
        double x = 0;
        double y = 0;
        
        x += A ? -10 : 0;
        x += D ? 10 : 0;
        y += W ? -10 : 0;
        y += (S && !LeftCtrl) ? 10 : 0;
        
        if (SelectedFigure != null)
        {
            SelectedFigure.Move(x, y);
            IsSaved = false;
        }
        else if (SelectedImage != null)
        {
            SelectedImage.Move(x, y);
            IsSaved = false;
        }

        if (LeftCtrl && S)
        {
            Save(Path);
        }
    }

    private async Task<bool> Save(string? path)
    {
        if (path == null)
        {
            var dialog = new SaveFileDialog
            {
                Title = "Save as",
                Filters = new List<FileDialogFilter>
                {
                    new FileDialogFilter { Name = "Oop_img Files", Extensions = new List<string> { "oop_img" } },
                },
                DefaultExtension = "oop_img",
                InitialFileName = "new_drawing.oop_img",
            };

            path = await dialog.ShowAsync(this);

            if (path == null)
            {
                return false;
            }
        }
        
        List<lib.Image> images = [];
        foreach (var imgControl in this.Main.Children)
        {
            var img = (imgControl as Canvas)?.Tag as lib.Image;
            if (img == null) { continue; }
            images.Add(img);
        }
        
        lib.Serializer.Write(images, path);

        Path = path;
        
        return true;
    }

    private async Task<bool> Open()
    {
        var dialog = new OpenFileDialog
        {
            Title = "Open File",
            Filters = new List<FileDialogFilter>
            {
                new FileDialogFilter { Name = "Oop_img Files", Extensions = new List<string> { "oop_img" } },
            },
            AllowMultiple = true,
        };

        var result = await dialog.ShowAsync(this);
        if (result != null && result.Length > 0)
        {
            foreach (string path in result)
            {
                foreach (lib.Image img in lib.Serializer.Read(path))
                {
                    Main.Children.Add(img.Canvas);
                }
            }

            if (result.Length == 1)
            {
                Path = result[0];
            }
            
            return true;
        }

        return false;
    }
    
    private async void SaveMenuItem_OnClick(object? sender, RoutedEventArgs e)
    {
        if (!(await Save(null)))
        {
            ShowError("Interrupted");
        }
    }
    
    private async void OpenMenuItem_OnClick(object? sender, RoutedEventArgs e)
    {
        if (!(await Open()))
        {
            ShowError("Interrupted");
        }
    }
    
    private async void ShowError(string errorText)
    {
        MessageRow.Text = errorText;
        await Tick(5000);
        MessageRow.Text = "";
    }
    
    async Task Tick(int delay)
    {
        await Task.Delay(delay);
    }

    private void Delete_OnClick(object? sender, RoutedEventArgs e)
    {
        MainWindow? mainWindow = null;
        if (SelectedFigure != null)
        {
            mainWindow = SelectedFigure.ParentImage?.Canvas.GetVisualRoot() as MainWindow;
            
            SelectedFigure.Destroy();
            SelectedFigure = null;
        } else if (SelectedImage != null)
        {
            mainWindow = SelectedImage.Canvas.GetVisualRoot() as MainWindow; 
            SelectedImage.Destroy();
            SelectedImage = null;
        }

        if (mainWindow != null)
        {
            mainWindow.EditPanel.Children.Clear();
            mainWindow.ApplyButton.IsVisible = false;
            mainWindow.DeleteButton.IsVisible = false;
        }
    }
}