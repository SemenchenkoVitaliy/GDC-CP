using System;
using System.Collections.Generic;

namespace KpiLex.Domain.Models
{
    public class Lecture
    {
        public Guid LectureId { get; set; }
        public string Title { get; set; }
        public string Description { get; set; }
        public int Likes { get; set; }
        public int Dislikes { get; set; }
        public Uri Uri { get; set; }
        public ICollection<string> Tags { get; set; }
    }
}