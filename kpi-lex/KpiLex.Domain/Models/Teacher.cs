using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class Teacher : User
    {
        public int TeacherId { get; set; }
        public string Position { get; set; }
        public Faculty Faculty { get; set; }
        public Speciality Speciality { get; set; }
        public ICollection<Course> Courses { get; set; }
    }
}