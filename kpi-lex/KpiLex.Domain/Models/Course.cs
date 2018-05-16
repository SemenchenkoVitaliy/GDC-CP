using System;
using System.Collections;
using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class Course
    {
        public long CourseId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public DateTime CreationDate { get; set; }
        public int Likes { get; set; }
        public int Dislikes { get; set; }
        public ICollection<Lecture> Lectures { get; set; }
        public ICollection<CourseComment> CourseComments { get; set; }
    }
}