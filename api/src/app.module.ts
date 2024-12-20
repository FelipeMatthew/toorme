import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsersModule } from './users/users.module';
import { TypeOrmModule } from "@nestjs/typeorm"

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'postgres',
      host: 'localhost',
      port: 5432,
      username: 'postgres',
      password: 'postgres',
      database: 'toorme',
      entities: [__dirname + "/../**/*.entity.{js,ts}"],
      synchronize: true, // todo: remove later
    })
    ,UsersModule],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
