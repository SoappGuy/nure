namespace pz3;

class Program
{
    static void Main(string[] args)
    {
        var comp1 = new Complex(3, -2);
        var comp2 = new Complex(-4, 3);
        Console.WriteLine(Complex.Mult(comp1, comp2));
        Console.WriteLine("a");
    }
}