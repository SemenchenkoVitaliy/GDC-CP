using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class Faculty
    {
        public int FacultyId { get; set; }
        public string Name { get; set; }
        public ICollection<Speciality> Specialities { get; set; }
    }
}