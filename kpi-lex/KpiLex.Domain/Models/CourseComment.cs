using System;

namespace KpiLex.Domain.Models
{
    public class CourseComment
    {
        public long CourseCommentId { get; set; }
        public string Content { get; set; }
        public DateTime CreationDate { get; set; }
    }
}