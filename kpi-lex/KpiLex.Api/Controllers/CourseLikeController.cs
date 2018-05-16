using System.Threading.Tasks;
using KpiLex.Api.Constants;
using KpiLex.Api.Mappers.Abstract;
using KpiLex.Api.Models.Response;
using KpiLex.BusinessLogic.Services.Abstract;
using KpiLex.Domain.Models;
using Microsoft.AspNetCore.Mvc;

namespace KpiLex.Api.Controllers
{
    [Route(RouteConstants.CourseLikeRoute)]
    public class CourseLikeController : Controller
    {
        private readonly ICourseService _courseService;
        private readonly IApiMapper<Course, CourseLikeResponseModel> _responseModelMapper;
        
        public CourseLikeController(ICourseService courseService,
            IApiMapper<Course, CourseLikeResponseModel> responseModelMapper)
        {
            _courseService = courseService;
            _responseModelMapper = responseModelMapper;
        }

        [HttpGet]
        public async Task<CourseLikeResponseModel> GetLikes([FromRoute] long courseId)
        {
            var course = await _courseService.GetById(courseId);
            return _responseModelMapper.Map(course);
        }
        
        [HttpPut, Route("{value}")]
        public async Task<IActionResult> PutLikes(
            [FromRoute] int courseId, 
            [FromRoute] int value)
        {
            await _courseService.AddLikeAsync(courseId, value);
            return Ok();
        }
    }
}