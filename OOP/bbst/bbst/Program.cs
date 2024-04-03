namespace bbst;

class Program
{
    static void Main(string[] args)
    {
        var bbst = new Bbst();
        
        bbst.Add(3);
        bbst.Add(5);
        bbst.Add(1);
        bbst.Add(-1);
        bbst.Add(2);
        
        bbst.PrettiePrint();
    }
}