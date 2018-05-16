namespace KpiLex.Api.Constants
{
    public class RouteConstants
    {
        public const string BaseRoute = "api/v1";
        public const string UserRoute = BaseRoute + "/users";
        public const string StudentRoute = BaseRoute + "/students";
        public const string TeacherRoute = BaseRoute + "/teachers";
        public const string LectureRoute = BaseRoute + "/lectures";
        public const string CourseRoute = BaseRoute + "/courses";
        public const string CourseLikeRoute = CourseRoute + "/{courseId}/likes";
        public const string CourseCommentRoute = CourseRoute + "/{courseId}/comments";
        public const string FacultyRoute = BaseRoute + "/faculties";
        public const string SpecialitiesRoute = BaseRoute + "/specialities";
        public const string StudentGroupRoute = BaseRoute + "/groups";
        public const string WatchLaterRoute = StudentRoute + "/{studentId}/watch_later";
    }
}