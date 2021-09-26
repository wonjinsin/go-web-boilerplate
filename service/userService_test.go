package service

import (
	"context"
	"pikachu/mock"
	"pikachu/model"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/juju/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("userService Test", func() {
	var (
		mockCtrl *gomock.Controller
		muRepo   *mock.MockUserRepository

		userService UserService
		ctx         context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		muRepo = mock.NewMockUserRepository(mockCtrl)

		userService = NewUserService(muRepo)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("NewUser", func() {
		var (
			email string
			nick  string

			mockedUser    *model.User
			expectedUser  *model.User
			returnedUser  *model.User
			returnedError error

			isError    error
			isErrorNil error
		)
		BeforeEach(func() {
			email = gofakeit.Email()
			nick = gofakeit.Name()

			mockedUser = &model.User{
				Email: email,
				Nick:  nick,
			}

			expectedUser = mockedUser
			isError = errors.AlreadyExistsf("User already exists")
			isErrorNil = nil
		})
		JustBeforeEach(func() {
			muRepo.EXPECT().
				NewUser(gomock.Any(), gomock.Eq(mockedUser)).
				Return(expectedUser, isErrorNil).AnyTimes()

			muRepo.EXPECT().
				GetUserByEmail(gomock.Any(), gomock.Eq(mockedUser.Email)).
				Return(isErrorNil, isError).AnyTimes()

			returnedUser, returnedError = userService.NewUser(ctx, mockedUser)
		})
		Context("normal", func() {
			It("should not error", func() {
				Expect(mockedUser.Email).To(Equal(returnedUser.Email))
				Expect(mockedUser.Nick).To(Equal(returnedUser.Nick))
				Expect(returnedError).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GetUser", func() {
		var (
			uid   string
			email string
			nick  string

			expectedUser  *model.User
			expectedError error
			returnedUser  *model.User
			returnedError error
		)
		BeforeEach(func() {
			uid = uuid.New().String()
			email = gofakeit.Email()
			nick = gofakeit.Name()

			expectedUser = &model.User{
				UID:   uid,
				Email: email,
				Nick:  nick,
			}
			expectedError = nil
		})
		JustBeforeEach(func() {
			muRepo.EXPECT().
				GetUser(gomock.Any(), gomock.Eq(uid)).
				Return(expectedUser, expectedError).AnyTimes()

			returnedUser, returnedError = userService.GetUser(ctx, uid)
		})
		Context("normal", func() {
			It("should not error", func() {
				Expect(expectedUser).To(Equal(returnedUser))
				Expect(returnedError).NotTo(HaveOccurred())
			})
		})
	})

	Describe("GetUserByEmail", func() {
		var (
			uid   string
			email string
			nick  string

			expectedUser  *model.User
			expectedError error
			returnedUser  *model.User
			returnedError error
		)
		BeforeEach(func() {
			uid = uuid.New().String()
			email = gofakeit.Email()
			nick = gofakeit.Name()

			expectedUser = &model.User{
				UID:   uid,
				Email: email,
				Nick:  nick,
			}
			expectedError = nil
		})
		JustBeforeEach(func() {
			muRepo.EXPECT().
				GetUserByEmail(gomock.Any(), gomock.Eq(email)).
				Return(expectedUser, expectedError).AnyTimes()

			returnedUser, returnedError = userService.GetUserByEmail(ctx, email)
		})
		Context("normal", func() {
			It("should not error", func() {
				Expect(expectedUser).To(Equal(returnedUser))
				Expect(returnedError).NotTo(HaveOccurred())
			})
		})
	})

	Describe("UpdateUser", func() {
		var (
			uid   string
			email string
			nick  string

			mockedUser    *model.User
			newUser       *model.User
			expectedUser  *model.User
			expectedError error
			returnedUser  *model.User
			returnedError error
		)
		BeforeEach(func() {
			uid = uuid.New().String()
			email = gofakeit.Email()
			nick = gofakeit.Name()

			mockedUser = &model.User{
				Email: email,
				Nick:  nick,
			}

			newUser = &model.User{
				UID:   uid,
				Email: email,
				Nick:  nick,
			}

			expectedUser = &model.User{
				UID:   uid,
				Email: email,
				Nick:  nick,
			}
			expectedError = nil
		})
		JustBeforeEach(func() {
			muRepo.EXPECT().
				GetUser(gomock.Any(), gomock.Eq(uid)).
				Return(newUser, expectedError).AnyTimes()

			muRepo.EXPECT().
				UpdateUser(gomock.Any(), gomock.Eq(newUser)).
				Return(expectedUser, expectedError).AnyTimes()

			returnedUser, returnedError = userService.UpdateUser(ctx, uid, mockedUser)
		})
		Context("normal", func() {
			It("should not error", func() {
				Expect(expectedUser).To(Equal(returnedUser))
				Expect(returnedError).NotTo(HaveOccurred())
			})
		})
	})

	Describe("DeleteUser", func() {
		var (
			uid string

			expectedError error
			returnedError error
		)
		BeforeEach(func() {
			uid = uuid.New().String()
			expectedError = nil
		})
		JustBeforeEach(func() {
			muRepo.EXPECT().
				DeleteUser(gomock.Any(), gomock.Eq(uid)).
				Return(expectedError).AnyTimes()

			returnedError = userService.DeleteUser(ctx, uid)
		})
		Context("normal", func() {
			It("should not error", func() {
				Expect(returnedError).NotTo(HaveOccurred())
			})
		})

	})

})
