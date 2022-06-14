package rc

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getStudentEnrollment(ctx *gin.Context) {
	rid := ctx.Param("rid")

	sid, err := getStudentRecruitmentCycleID(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	var questions []RecruitmentCycleQuestion
	var answers []RecruitmentCycleQuestionsAnswer

	err = fetchStudentQuestions(ctx, rid, &questions)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = fetchStudentAnswers(ctx, strconv.FormatUint(uint64(sid), 10), &answers)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"questions": questions, "answers": answers})
}

func postEnrollmentAnswer(ctx *gin.Context) {
	rid := ctx.Param("rid")
	var answer RecruitmentCycleQuestionsAnswer

	err := ctx.ShouldBindJSON(&answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	answer.StudentRecruitmentCycleID, err = getStudentRecruitmentCycleID(ctx, rid)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = createStudentAnswer(ctx, &answer)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": fmt.Sprintf("Answer %d created", answer.ID)})
}
