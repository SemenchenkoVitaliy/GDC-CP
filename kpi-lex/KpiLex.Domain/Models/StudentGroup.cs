using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class StudentGroup
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public Speciality Speciality { get; set; }
    }
}