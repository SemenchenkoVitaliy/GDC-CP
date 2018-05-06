using System.Collections.Generic;
using System.Threading.Tasks;
using KpiLex.Domain.Models;

namespace KpiLex.BusinessLogic.Services.Abstract
{
    public interface ICourseCommentService : IService<CourseComment, long>
    {
        Task<IReadOnlyCollection<CourseComment>> GetCourseComments(long courseId);
    }
}