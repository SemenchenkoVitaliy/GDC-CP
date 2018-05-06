using System.Text.RegularExpressions;

namespace KpiLex.Domain.Models
{
    public class Student : User
    {
        public int StudentId { get; set; }
        public Speciality Speciality { get; set; }
        public StudentGroup StudentGroup { get; set; }
    }
}