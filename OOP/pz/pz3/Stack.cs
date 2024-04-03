namespace pz3;

public class Stack
{
    private int[] _arr;
    public int Top { get; private set; } = 0;

    public Stack()
    {
        this._arr = new int[10];
    }
    
    public Stack(int cap)
    {
        this._arr = new int[cap];
    }

    public void Push(int data)
    {
        if (this.Top == this._arr.Length)
        {
            Console.WriteLine("Full");
            return;
        }

        this._arr[this.Top++] = data;

    }
    public int Pop()
    {
        int data;
        
        if (this.Top == 0)
        {
            Console.WriteLine("Empty");
            return -1;
        }
        return this._arr[--this.Top];
    }

    public override string ToString()
    {
        if (this.Top == 0)
        {
            return "Empty";
        }
        
        string result = "";
        for (int i = 0; i < this.Top; i++)
        {
            result += this._arr[i] + ", ";
        }

        return result[..^2];

    }
}