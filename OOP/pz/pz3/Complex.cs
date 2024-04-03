namespace pz3;

public class Complex
{
    private int real;
    private int imagen;

    public Complex(int real, int imagen)
    {
        this.real = real;
        this.imagen = imagen;
    }

    public static Complex Add(Complex self, Complex other)
    {
        int real = self.real + other.real;
        int imagen = self.imagen + other.imagen;
        
        return new Complex(real, imagen);
    }
    
    public static Complex Mult(Complex self, Complex other)
    {
        int real = self.real * other.real - self.imagen * other.imagen;
        int imagen = self.real * other.imagen + self.imagen * other.real;
        
        return new Complex(real, imagen);
    }
    
    public override string ToString()
    {
        char sym;
        if (this.imagen < 0)
        {
            sym = '-';
        }
        else
        {
            sym = '+';
        }
        
        return $"{this.real} {sym} {Math.Abs(this.imagen)}i";
    }
}