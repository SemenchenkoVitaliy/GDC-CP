using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class Speciality
    {
        public int SpecialityId { get; set; }
        public string Name { get; set; }
        public ICollection<Teacher> Teachers { get; set; }
    }
}