using System;
using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class WatchLaterItem
    {
        public Guid WatchLaterId { get; set; }
        public Student Student { get; set; }
        public ICollection<Course> Courses { get; set; }
        public ICollection<Lecture> Lectures { get; set; }
    }
}