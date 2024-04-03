namespace bbst;

public class Node
{
    public int Data;
    public Node? Left = null;
    public Node? Right = null;
    private int height = 1;

    public Node(int data)
    {
        this.Data = data;
    }

    public bool Add(int data)
    {
        bool result;
        if (data == this.Data)
        {
            result = false;
        } 
        else if (data < this.Data)
        {
            if (this.Left != null)
            {
                result = this.Left.Add(data);
            }
            else
            {
                this.Left = new Node(data);
                result = true;
            }
        }
        else
        {
            if (this.Right != null)
            {
                result = this.Right.Add(data);
            }
            else
            {
                this.Right = new Node(data);
                result = true;
            }
        }

        if (result)
        {
            int leftHeight = this.Left?.height ?? 0;
            int rightHeight = this.Right?.height ?? 0;

            this.height = Math.Max(leftHeight, rightHeight) + 1;
        }

        return result;
    }
    
    public void PrettiePrint(string padding, string pointer, bool top)
    {
        if (this.Right != null)
        {
            this.Right.PrettiePrint(padding + (top ? "\u2502   " : "    "), "\u250c\u2500\u2500", false);
        }

        Console.WriteLine($"{padding}{pointer}({this.Data})");

        if (this.Left != null)
        {
            this.Left.PrettiePrint(padding + (!top ? "\u2502   " : "    "), "\u2514\u2500\u2500", true);
        }

    }
}

public class Bbst
{
    private Node? Head = null;
    public int Size { get; private set; } = 0;

    public void Add(int data)
    {
        if (this.Head != null)
        {
            if (this.Head.Add(data))
            {
                this.Size++;
            }
            else
            {
                Console.WriteLine("Element already in bst, skipping");
            }
        }
        else
        {
            this.Head = new Node(data);
            this.Size++;
        }
    }
    
    public void PrettiePrint()
    {
        if (this.Head != null)
        {
            if (this.Head.Right != null)
            {
                this.Head.Right.PrettiePrint("    ", "\u250c\u2500\u2500", false);
            }

            Console.WriteLine($"->({this.Head.Data})");

            if (this.Head.Left != null)
            {
                this.Head.Left.PrettiePrint("    ", "\u2514\u2500\u2500", true);
            }

            Console.WriteLine($"size: {this.Size}");
        }
        else
        {
            Console.WriteLine("Empty");
        }
    } 
}