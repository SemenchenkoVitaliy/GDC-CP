using System.Collections.Generic;
using System.Threading.Tasks;
using KpiLex.Domain.Models;

namespace KpiLex.BusinessLogic.Services.Abstract
{
    public interface ICourseService : IService<Course, long>
    {
        Task AddLikeAsync(int courseId, int value);
        Task<IReadOnlyCollection<Course>> GetCoursesByTags(IReadOnlyCollection<string> tags);
    }
}