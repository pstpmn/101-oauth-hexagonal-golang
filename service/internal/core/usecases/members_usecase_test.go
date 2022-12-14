package usecases

import (
	"errors"
	domain "learn-oauth2/internal/core/domains"
	ports2 "learn-oauth2/internal/core/ports"
	"learn-oauth2/internal/core/ports/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

func Test_membersUseCase_NewMember(t *testing.T) {
	uuid := "f43ab0cc-8653-42dc-853d-fdee58a17cd6"
	mockRequest := &domain.Members{Mid: uuid, Username: "1", Password: "root", FirstName: "root", LastName: "root", DateOfBird: time.Now(), RegisterType: 1, CreatedAt: time.Now()}
	mockMembersRepo := new(mocks.MembersRepository)
	mockMembersRepoCaseTwo := new(mocks.MembersRepository)
	mockCatepgoriesRepo := new(mocks.RegisterCategories)
	mockUidService := new(mocks.IUuidService)
	mockCryptoService := new(mocks.ICryptoService)

	mockMembersRepo.On("Create", mock.AnythingOfType("*domains.Members")).Return(mockRequest, nil)
	mockMembersRepo.On("GetByUser", mock.AnythingOfType("string")).Return(&domain.Members{}, nil)
	mockUidService.On("Random").Return(uuid)
	mockCryptoService.On("Bcrypt", mock.AnythingOfType("string")).Return("encript", nil)
	mockMembersRepoCaseTwo.On("Create", mock.AnythingOfType("*domains.Members")).Return(&domain.Members{}, nil)
	mockMembersRepoCaseTwo.On("GetByUser", mock.AnythingOfType("string")).Return(mockRequest, nil)

	type fields struct {
		membersRepo        ports2.MembersRepository
		RegisterCategories ports2.RegisterCategories
		uidService         ports2.IUuidService
		cryptoService      ports2.ICryptoService
	}
	type args struct {
		user     string
		pass     string
		fistName string
		lastName string
		dob      time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Members
		wantErr bool
	}{
		{
			"test create member should be success",
			fields{
				mockMembersRepo,
				mockCatepgoriesRepo,
				mockUidService,
				mockCryptoService,
			},
			args{
				"root",
				"root",
				"root",
				"root",
				time.Now(),
			},
			mockRequest,
			false,
		},
		{
			"test create member should be error because username is used",
			fields{
				mockMembersRepoCaseTwo,
				mockCatepgoriesRepo,
				mockUidService,
				mockCryptoService,
			},
			args{
				"root",
				"root",
				"root",
				"root",
				time.Now(),
			},
			&domain.Members{},
			true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			m := membersUseCase{
				membersRepo:        tt.fields.membersRepo,
				RegisterCategories: tt.fields.RegisterCategories,
				UidService:         tt.fields.uidService,
				CryptoService:      tt.fields.cryptoService,
			}
			got, err := m.NewMember(tt.args.user, tt.args.pass, tt.args.fistName, tt.args.lastName, tt.args.dob)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMember() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_membersUseCase_FindMemberById(t *testing.T) {
	id := "9cb6eb12-0b83-4034-92de-61d0e6256699"
	mockResponse := &domain.Members{Mid: id, Username: "root", Password: "root", FirstName: "root", LastName: "root", DateOfBird: time.Now(), RegisterType: 1, CreatedAt: time.Now()}
	mockMembersRepoCaseOne := new(mocks.MembersRepository)
	mockMembersRepoCaseTwo := new(mocks.MembersRepository)

	mockCatepgoriesRepo := new(mocks.RegisterCategories)
	mockMembersRepoCaseOne.On("Get", mock.AnythingOfType("string")).Return(mockResponse, nil)
	mockMembersRepoCaseTwo.On("Get", "random").Return(&domain.Members{}, nil)

	type fields struct {
		membersRepo        ports2.MembersRepository
		RegisterCategories ports2.RegisterCategories
	}

	type args struct {
		id string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Members
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test get member by id should success",
			fields{mockMembersRepoCaseOne, mockCatepgoriesRepo},
			args{id: id},
			mockResponse,
			false,
		},
		{
			"test get member by id should error because not found member",
			fields{mockMembersRepoCaseTwo, mockCatepgoriesRepo},
			args{id: "random"},
			&domain.Members{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersUseCase{
				membersRepo:        tt.fields.membersRepo,
				RegisterCategories: tt.fields.RegisterCategories,
			}
			got, err := m.FindMemberById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMemberById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMemberById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_membersUseCase_Authentication(t *testing.T) {
	uuid := "f43ab0cc-8653-42dc-853d-fdee58a17cd6"
	key := "secret"

	mockResponse := &domain.Members{Mid: uuid, Username: "root", Password: "root", FirstName: "root", LastName: "root", DateOfBird: time.Now(), RegisterType: 1, CreatedAt: time.Now()}
	mockMembersRepo := new(mocks.MembersRepository)
	mockMembersRepoCaseTwo := new(mocks.MembersRepository)
	mockCatepgoriesRepo := new(mocks.RegisterCategories)
	mockUidService := new(mocks.IUuidService)
	mockCryptoService := new(mocks.ICryptoService)
	mockJwtService := new(mocks.IJwtService)

	mockMembersRepo.On("GetByUser", mock.AnythingOfType("string")).Return(mockResponse)
	mockMembersRepoCaseTwo.On("GetByUser", mock.AnythingOfType("string")).Return(&domain.Members{})
	//mockUidService.On("Random").Return(uuid)
	mockCryptoService.On("ValidateBcrypt", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true)
	mockJwtService.On("Generate", mock.AnythingOfType("map[string]interface {}"), "secret", mock.AnythingOfType("time.Time")).Return("mockToken", nil)

	type fields struct {
		membersRepo        ports2.MembersRepository
		RegisterCategories ports2.RegisterCategories
		UidService         ports2.IUuidService
		CryptoService      ports2.ICryptoService
		JwtService         ports2.IJwtService
	}
	type args struct {
		user string
		pass string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Members
		wantErr bool
	}{
		{
			"test usecase should be auth success",
			fields{mockMembersRepo, mockCatepgoriesRepo, mockUidService, mockCryptoService, mockJwtService},
			args{
				"root",
				"root",
			},
			mockResponse,
			false,
		},
		{
			"test usecase should be fail because username not found",
			fields{mockMembersRepoCaseTwo, mockCatepgoriesRepo, mockUidService, mockCryptoService, mockJwtService},
			args{
				"root",
				"root",
			},
			&domain.Members{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersUseCase{
				membersRepo:        tt.fields.membersRepo,
				RegisterCategories: tt.fields.RegisterCategories,
				UidService:         tt.fields.UidService,
				CryptoService:      tt.fields.CryptoService,
				JwtService:         mockJwtService,
			}

			_, got, err := m.Authentication(tt.args.user, tt.args.pass, key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authentication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authentication() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_membersUseCase_Authorization(t *testing.T) {
	token := "testToken"
	key := "secret"

	mockMembersRepo := new(mocks.MembersRepository)
	mockCatepgoriesRepo := new(mocks.RegisterCategories)
	mockUidService := new(mocks.IUuidService)
	mockCryptoService := new(mocks.ICryptoService)
	mockJwtService := new(mocks.IJwtService)
	mockJwtServiceCaseTwo := new(mocks.IJwtService)

	mockJwtService.On("Extract", token, key).Return(map[string]interface{}{}, nil)
	mockJwtServiceCaseTwo.On("Extract", token, key).Return(map[string]interface{}{}, errors.New("error invalid token"))

	type fields struct {
		membersRepo        ports2.MembersRepository
		RegisterCategories ports2.RegisterCategories
		UidService         ports2.IUuidService
		CryptoService      ports2.ICryptoService
		JwtService         ports2.IJwtService
	}
	type args struct {
		token string
		key   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want    map[string]string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test authorization token should be success",
			fields{
				membersRepo:        mockMembersRepo,
				RegisterCategories: mockCatepgoriesRepo,
				UidService:         mockUidService,
				CryptoService:      mockCryptoService,
				JwtService:         mockJwtService,
			},
			args{
				token: token,
				key:   key,
			},
			//nil,
			false,
		},
		{
			"test authorization invalid token should be fails",
			fields{
				membersRepo:        mockMembersRepo,
				RegisterCategories: mockCatepgoriesRepo,
				UidService:         mockUidService,
				CryptoService:      mockCryptoService,
				JwtService:         mockJwtServiceCaseTwo,
			},
			args{
				token: token,
				key:   key,
			},
			//nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := membersUseCase{
				membersRepo:        tt.fields.membersRepo,
				RegisterCategories: tt.fields.RegisterCategories,
				UidService:         tt.fields.UidService,
				CryptoService:      tt.fields.CryptoService,
				JwtService:         tt.fields.JwtService,
			}
			_, err := m.Authorization(tt.args.token, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Authorization() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
